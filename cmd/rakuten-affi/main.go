package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/kohge4/go-rakutenapi/rakuten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/masa23/rakuten-affi/cmd/rakuten-affi/config"
)

var conf *config.Config

func rakutenSearch(keyword string) (*rakuten.IchibaItemResponse, error) {
	ctx := context.Background()
	tp := rakuten.Transport{}

	client := rakuten.NewClient(tp.Client(), conf.ApplicationID, conf.AffiliateID)

	// QueryParameter for Search argument
	sOptions := &rakuten.IchibaItemSearchParams{
		Keyword: keyword,
		Hits:    20,
	}

	// Search Items from Rakuten Ichiba API
	ichiba, _, err := client.Ichiba.Search(ctx, sOptions)
	if err != nil {
		return nil, err
	}

	return ichiba, nil
}

func search(c echo.Context) error {
	keyword := c.FormValue("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, "keyword is required")
	}

	result, err := rakutenSearch(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func main() {
	var confPath string
	var listenPort string
	var err error
	flag.StringVar(&confPath, "conf", "config.yaml", "path to config file")
	flag.StringVar(&listenPort, "port", ":8080", "port to listen")
	flag.Parse()

	conf, err = config.Load(confPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// echo API
	e := echo.New()

	// CORSの許可
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/search", search)

	e.Logger.Fatal(e.Start(listenPort))
}
