package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"

	paapi5 "github.com/goark/pa-api"
	"github.com/goark/pa-api/entity"
	"github.com/goark/pa-api/query"
	"github.com/kohge4/go-rakutenapi/rakuten"
	"github.com/labstack/echo/v4"
	"github.com/masa23/search-api/cmd/search-api/config"
)

var conf *config.Config

type SortType string

const (
	SortTypeDefault   SortType = "default"
	SortTypePriceAsc  SortType = "price_asc"
	SortTypePriceDesc SortType = "price_desc"
)

type SearchRequest struct {
	Keyword  string   `json:"keyword"`
	MinPrice int      `json:"minPrice"`
	MaxPrice int      `json:"maxPrice"`
	Sort     SortType `json:"sort"`
}

func (s *SearchRequest) validate() error {
	if s.Keyword == "" {
		return errors.New("keyword is required")
	}
	if s.MinPrice < 0 {
		return errors.New("minPrice must be greater than or equal to 0")
	}
	if s.MaxPrice < 0 {
		return errors.New("maxPrice must be greater than or equal to 0")
	}
	if s.MinPrice > s.MaxPrice {
		return errors.New("minPrice must be less than or equal to maxPrice")
	}
	if s.Sort == "" {
		s.Sort = SortTypeDefault
	}
	if s.Sort != SortTypeDefault && s.Sort != SortTypePriceAsc && s.Sort != SortTypePriceDesc {
		return errors.New("sort is invalid")
	}
	return nil
}

func amazonItemSearch(s *SearchRequest) (interface{}, error) {
	client := paapi5.New(
		paapi5.WithMarketplace(paapi5.LocaleJapan),
	).CreateClient(
		conf.Amazon.AssociateTag,
		conf.Amazon.AccessKey,
		conf.Amazon.SecretKey,
	)

	filterMap := make(query.RequestMap)
	filterMap[query.CurrencyOfPreference] = "JPY"

	if s.MinPrice > 0 {
		filterMap[query.MinPrice] = s.MinPrice * 100
	}
	if s.MaxPrice > 0 {
		filterMap[query.MaxPrice] = s.MaxPrice * 100
	}
	// Price:HighToLow", "Price:LowToHigh"
	if s.Sort == SortTypePriceAsc {
		filterMap[query.SortBy] = "Price:LowToHigh"
	}
	if s.Sort == SortTypePriceDesc {
		filterMap[query.SortBy] = "Price:HighToLow"
	}

	q := query.NewSearchItems(
		client.Marketplace(),
		client.PartnerTag(),
		client.PartnerType(),
	).Search(query.Keywords, s.Keyword).EnableImages().EnableItemInfo().EnableOffers().RequestFilters(filterMap)

	body, err := client.RequestContext(context.Background(), q)
	if err != nil {
		return nil, err
	}

	res, err := entity.DecodeResponse(body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func amazonSearch(c echo.Context) error {
	searchRequest := new(SearchRequest)
	if err := c.Bind(searchRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := searchRequest.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := amazonItemSearch(searchRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func rakutenItemSearch(s *SearchRequest) (*rakuten.IchibaItemResponse, error) {
	ctx := context.Background()
	tp := rakuten.Transport{}

	client := rakuten.NewClient(tp.Client(), conf.Rakuten.ApplicationID, conf.Rakuten.AffiliateID)

	sOptions := &rakuten.IchibaItemSearchParams{
		Keyword:  s.Keyword,
		MinPrice: s.MinPrice,
		MaxPrice: s.MaxPrice,
		Hits:     10,
	}

	if s.Sort == SortTypePriceAsc {
		sOptions.Sort = "+itemPrice"
	}
	if s.Sort == SortTypePriceDesc {
		sOptions.Sort = "-itemPrice"
	}

	ichiba, _, err := client.Ichiba.Search(ctx, sOptions)
	if err != nil {
		return nil, err
	}

	return ichiba, nil
}

func rakutenSearch(c echo.Context) error {
	searchRequest := new(SearchRequest)
	if err := c.Bind(searchRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := searchRequest.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := rakutenItemSearch(searchRequest)
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
	/*e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))*/

	e.POST("/amazon/search", amazonSearch)
	e.POST("/rakuten/search", rakutenSearch)

	e.Logger.Fatal(e.Start(listenPort))
}
