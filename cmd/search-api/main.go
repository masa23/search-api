package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	paapi5 "github.com/goark/pa-api"
	"github.com/goark/pa-api/entity"
	"github.com/goark/pa-api/query"
	"github.com/kohge4/go-rakutenapi/rakuten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/masa23/search-api/cmd/search-api/config"
)

var conf *config.Config

func amazonItemSearch(keyword string) (interface{}, error) {
	client := paapi5.New(
		paapi5.WithMarketplace(paapi5.LocaleJapan),
	).CreateClient(
		conf.Amazon.AssociateTag,
		conf.Amazon.AccessKey,
		conf.Amazon.SecretKey,
	)

	q := query.NewSearchItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).Search(query.Keywords, keyword).EnableImages().EnableItemInfo().EnableOffers()
	//Request and response
	body, err := client.RequestContext(context.Background(), q)
	if err != nil {
		return nil, err
	}
	//io.Copy(os.Stdout, bytes.NewReader(body))

	//Decode JSON
	res, err := entity.DecodeResponse(body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func amazonSearch(c echo.Context) error {
	keyword := c.FormValue("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, "keyword is required")
	}

	result, err := amazonItemSearch(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func rakutenItemSearch(keyword string) (*rakuten.IchibaItemResponse, error) {
	ctx := context.Background()
	tp := rakuten.Transport{}

	client := rakuten.NewClient(tp.Client(), conf.Rakuten.ApplicationID, conf.Rakuten.AffiliateID)

	// QueryParameter for Search argument
	sOptions := &rakuten.IchibaItemSearchParams{
		Keyword: keyword,
		Hits:    10,
	}

	// Search Items from Rakuten Ichiba API
	ichiba, _, err := client.Ichiba.Search(ctx, sOptions)
	if err != nil {
		return nil, err
	}

	return ichiba, nil
}

func rakutenSearch(c echo.Context) error {
	keyword := c.FormValue("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, "keyword is required")
	}

	result, err := rakutenItemSearch(keyword)
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

	e.POST("/amazon/search", amazonSearch)
	e.POST("/rakuten/search", rakutenSearch)

	e.Logger.Fatal(e.Start(listenPort))
}
