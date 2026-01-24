package rakuten

import (
	"context"
	"fmt"
)

// https://app.rakuten.co.jp/services/api/Product/Search/20170426?affiliateId=1a11cd3a.d53d02bd.1a11cd3b.54ce7249&applicationId=1058291654304154113&keyword=suchmos

// https://app.rakuten.co.jp/services/api/Product/Search/20170426?genreId=&genreInformationFlag=0&hits=0&keyword=suchmos&maxPrice=0&minPrice=0&orFlag=0&page=0&productId=&affiliateId=1a11cd3a.d53d02bd.1a11cd3b.54ce7249&applicationId=1058291654304154113

// parameter 設定の問題でエラー => min Price >1 , page >0

type IchibaProductSearchParams struct {
	Keyword              string `url:"keyword,omitempty"`
	GenreID              string `url:"genreId,omitempty"`
	ProductID            string `url:"productId,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	MinPrice             int    `url:"minPrice,omitempty"`
	MaxPrice             int    `url:"maxPrice,omitempty"`
	OrFlag               int    `url:"orFlag,omitempty"`
	GenreInformationFlag int    `url:"genreInformationFlag,omitempty"`
}

type IchibaProductResponse struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	PageCount int `json:"pageCount"`
	Products  []struct {
		Product struct {
			ProductID                 string  `json:"productId"`
			ProductName               string  `json:"productName"`
			ProductNo                 string  `json:"productNo"`
			BrandName                 string  `json:"brandName"`
			ProductURLPC              string  `json:"productUrlPC"`
			ProductURLMobile          string  `json:"productUrlMobile"`
			AffiliateURL              string  `json:"affiliateUrl"`
			SmallImageURL             string  `json:"smallImageUrl"`
			MediumImageURL            string  `json:"mediumImageUrl"`
			ProductCaption            string  `json:"productCaption"`
			ReleaseDate               string  `json:"releaseDate"`
			MakerCode                 string  `json:"makerCode"`
			MakerName                 string  `json:"makerName"`
			MakerNameKana             string  `json:"makerNameKana"`
			MakerNameFormal           string  `json:"makerNameFormal"`
			MakerPageURLPC            string  `json:"makerPageUrlPC"`
			MakerPageURLMobile        string  `json:"makerPageUrlMobile"`
			ItemCount                 int     `json:"itemCount"`
			SalesItemCount            int     `json:"salesItemCount"`
			UsedExcludeCount          int     `json:"usedExcludeCount"`
			UsedExcludeSalesItemCount int     `json:"usedExcludeSalesItemCount"`
			MaxPrice                  int     `json:"maxPrice"`
			SalesMaxPrice             int     `json:"salesMaxPrice"`
			UsedExcludeMaxPrice       int     `json:"usedExcludeMaxPrice"`
			UsedExcludeSalesMaxPrice  int     `json:"usedExcludeSalesMaxPrice"`
			MinPrice                  int     `json:"minPrice"`
			SalesMinPrice             int     `json:"salesMinPrice"`
			UsedExcludeMinPrice       int     `json:"usedExcludeMinPrice"`
			UsedExcludeSalesMinPrice  int     `json:"usedExcludeSalesMinPrice"`
			AveragePrice              int     `json:"averagePrice"`
			ReviewCount               int     `json:"reviewCount"`
			ReviewAverage             float64 `json:"reviewAverage"`
			ReviewURLPC               string  `json:"reviewUrlPC"`
			ReviewURLMobile           string  `json:"reviewUrlMobile"`
			Rank                      int     `json:"rank"`
			RankTargetGenreID         string  `json:"rankTargetGenreId"`
			RankTargetProductCount    int     `json:"rankTargetProductCount"`
			GenreID                   string  `json:"genreId"`
			GenreName                 string  `json:"genreName"`
			ProductDetails            []struct {
				Detail struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"detail"`
			} `json:"ProductDetails"`
		} `json:"Product"`
	} `json:"Products"`
	GenreInformation struct {
		Parent   []interface{} `json:"parent"`
		Current  []interface{} `json:"current"`
		Children []interface{} `json:"children"`
	} `json:"GenreInformation"`
}

func (s *IchibaService) ProductSearch(ctx context.Context, opt *IchibaProductSearchParams) (*IchibaProductResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Product/Search/20170426")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	ichibaProductResp := &IchibaProductResponse{}
	resp, err := s.client.Do(ctx, req, ichibaProductResp)
	if err != nil {
		return nil, resp, err
	}

	return ichibaProductResp, resp, nil
}
