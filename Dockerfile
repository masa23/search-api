# Multi-stage build
FROM node:20-alpine AS frontend-builder

WORKDIR /app

# Copy frontend package files
COPY frontend/package*.json ./

# Install frontend dependencies
RUN npm ci

# Copy frontend source code
COPY frontend/ .

# Build frontend
RUN npm run build-only

# Backend build stage
FROM golang:1.23-alpine AS backend-builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o search-api ./cmd/search-api

# Final stage
FROM alpine:latest

# Install nginx
RUN apk add --no-cache nginx

# Create nginx directories
RUN mkdir -p /run/nginx

# Remove default nginx config
RUN rm -rf /etc/nginx/conf.d/*

# Add nginx config
RUN echo " \
events { \
    worker_connections 1024; \
} \
\
http { \
    include       /etc/nginx/mime.types; \
    default_type  application/octet-stream; \
    sendfile      on; \
    tcp_nopush    on; \
    tcp_nodelay   on; \
    keepalive_timeout 65; \
    types_hash_max_size 2048; \
    \
    server { \
        listen 8081; \
        server_name localhost; \
        \
        location / { \
            root /usr/share/nginx/html; \
            index index.html; \
            try_files \$uri \$uri/ /index.html; \
        } \
        \
        location /api/ { \
            proxy_pass http://localhost:8080/; \
            proxy_set_header Host \$host; \
            proxy_set_header X-Real-IP \$remote_addr; \
            proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for; \
            proxy_set_header X-Forwarded-Proto \$scheme; \
        } \
    } \
} \
" > /etc/nginx/nginx.conf

# Create static files directory
RUN mkdir -p /usr/share/nginx/html

WORKDIR /root/

# Copy the backend binary from backend-builder stage
COPY --from=backend-builder /app/search-api .

# Copy the frontend build from frontend-builder stage
COPY --from=frontend-builder /app/dist /usr/share/nginx/html

# Copy the config file
COPY --from=backend-builder /app/cmd/search-api/config.sample.yaml ./config.yaml

# Expose port
EXPOSE 8081

# Create startup script
RUN echo '#!/bin/sh' > start.sh && \
    echo './search-api -port=:8080 &' >> start.sh && \
    echo 'nginx -g "daemon off;"' >> start.sh && \
    chmod +x start.sh

# Command to run the startup script
CMD ["./start.sh"]
