package rakuten

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
	"golang.org/x/oauth2"
)

const (
	defaultBaseURL   = "https://app.rakuten.co.jp/services/api/"
	defaultUserAgent = "rws"
)

// test URL = https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706?applicationId=1058291654304154113&keyword=suchmos

// Client is
type Client struct {
	client        *http.Client
	BaseURL       *url.URL
	UserAgent     string
	ApplicationID string
	AffiliateID   string
	Authenticator Authenticator
	OAuthToken    *oauth2.Token

	common service

	Ichiba   *IchibaService
	Books    *BooksService
	Travel   *TravelService
	Favorite *FavoService
	Recipe   *RecipeService
	Kobo     *KoboService
	Gora     *GoraService
}

type service struct {
	client *Client
}

type DefaultOptions struct {
	ApplicationID     string `url:"applicationId,omitempty"`
	AffiliateID       string `url:"affiliateId,omitempty"`
	ApplicationSecret string `url:"applicationSecret,omitempty"`
}

type DefaultTokenOptions struct {
	AccessToken string `url:"access_token,omitempty"`
	Format      string `url:"format,omitempty"`
}

// NewClient is
func NewClient(httpClient *http.Client, applicationID string, affiliateID string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:        httpClient,
		UserAgent:     defaultUserAgent,
		BaseURL:       baseURL,
		ApplicationID: applicationID,
		AffiliateID:   affiliateID,
	}
	c.common.client = c
	c.Ichiba = (*IchibaService)(&c.common)
	c.Favorite = (*FavoService)(&c.common)
	c.Books = (*BooksService)(&c.common)
	c.Travel = (*TravelService)(&c.common)
	c.Recipe = (*RecipeService)(&c.common)
	c.Kobo = (*KoboService)(&c.common)
	c.Gora = (*GoraService)(&c.common)

	return c
}

// NewRequest creates an API request.
// Param の 型　関連のエラーは　綺麗に書きたい
func (c *Client) NewRequest(method, urlSuffix string, urlParam interface{}, body interface{}) (*http.Request, error) {
	var dOpt interface{}

	if c.OAuthToken != nil {
		dOpt = DefaultTokenOptions{
			//AccessToken: c.OAuthToken.AccessToken,
			//Format:      "json",
		}
		fmt.Println("nilではない")
	} else {
		dOpt = DefaultOptions{
			ApplicationID: c.ApplicationID,
			AffiliateID:   c.AffiliateID,
		}
	}
	defaultParams, _ := query.Values(dOpt)

	urlStr, _ := addOptions(urlSuffix, urlParam)
	URL := urlStr + "&" + defaultParams.Encode()

	rel, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Response is a Apple Music API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	// 他のコードで Response 型 を 戻り値の型に使える
	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	// Response Body の処理
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}
	return response, err
}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt) //url.Values 型 https://github.com/google/go-querystring/blob/master/query/encode.go
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// Source represents the source of an error.
type Source struct {
	Parameter string      `json:"parameter"`
	Pointer   interface{} `json:"pointer"` // JSON pointer, A pointer to the associated entry in the request document.
}

// Error contains information about an error that occurred while processing a request.
type Error struct {
	ID     string      `json:"id"`
	About  string      `json:"about"`
	Status string      `json:"status"`
	Code   string      `json:"code"`
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
	Source Source      `json:"source"`
	Meta   interface{} `json:"meta"`
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Errors   []Error `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Errors)
}

type errorMessageResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

func (e *errorMessageResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		e.Message)
}

// UnauthorizedError occurs when server denied the request.
type UnauthorizedError errorMessageResponse

func (e *UnauthorizedError) Error() string {
	return (*errorMessageResponse)(e).Error()
}

// TooManyRequestsError represents the Too Many Requests (429) error when the server exceeds its capacity.
type TooManyRequestsError errorMessageResponse

func (e *TooManyRequestsError) Error() string {
	return (*errorMessageResponse)(e).Error()
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := ioutil.ReadAll(r.Body)

	switch r.StatusCode {
	case http.StatusUnauthorized:
		return &UnauthorizedError{
			Response: r,
			Message:  string(data),
		}
	case http.StatusTooManyRequests:
		errorMessageResponse := &TooManyRequestsError{Response: r}
		if err == nil && data != nil {
			json.Unmarshal(data, errorMessageResponse)
		}
		return errorMessageResponse
	default:
		errorResponse := &ErrorResponse{Response: r}
		if err == nil && data != nil {
			json.Unmarshal(data, errorResponse)
		}
		return errorResponse
	}
}

type Transport struct {
	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req) // per RoundTrip contract

	res, err := t.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	log.Printf("%d\t%s\t%s\n", res.StatusCode, req.Method, req.URL.String())

	return res, nil

}

// Client returns an *http.Client that makes requests.
func (t *Transport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
