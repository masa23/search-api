package rakuten

import (
	"context"
	"fmt"
)

type IchibaRankingParams struct {
	GenreID int    `url:"genreId,omitempty"`
	Age     int    `url:"age,omitempty"`
	Sex     int    `url:"sex,omitempty"`
	Carrier int    `url:"carrier,omitempty"`
	Page    int    `url:"page,omitempty"`
	Period  string `url:"period,omitempty"`
}

type IchibaRankingResponse struct {
	Items []struct {
		Item struct {
			MediumImageUrls []struct {
				ImageURL string `json:"imageUrl"`
			} `json:"mediumImageUrls"`
			PointRate         int    `json:"pointRate"`
			ShopOfTheYearFlag int    `json:"shopOfTheYearFlag"`
			AffiliateRate     string `json:"affiliateRate"`
			ShipOverseasFlag  int    `json:"shipOverseasFlag"`
			AsurakuFlag       int    `json:"asurakuFlag"`
			EndTime           string `json:"endTime"`
			TaxFlag           int    `json:"taxFlag"`
			StartTime         string `json:"startTime"`
			Rank              int    `json:"rank"`
			ItemCaption       string `json:"itemCaption"`
			Catchcopy         string `json:"catchcopy"`
			SmallImageUrls    []struct {
				ImageURL string `json:"imageUrl"`
			} `json:"smallImageUrls"`
			AsurakuClosingTime string `json:"asurakuClosingTime"`
			Carrier            int    `json:"carrier"`
			ImageFlag          int    `json:"imageFlag"`
			ShopAffiliateURL   string `json:"shopAffiliateUrl"`
			Availability       int    `json:"availability"`
			ItemCode           string `json:"itemCode"`
			PostageFlag        int    `json:"postageFlag"`
			ItemName           string `json:"itemName"`
			ItemPrice          string `json:"itemPrice"`
			PointRateEndTime   string `json:"pointRateEndTime"`
			ShopCode           string `json:"shopCode"`
			AffiliateURL       string `json:"affiliateUrl"`
			ShopName           string `json:"shopName"`
			AsurakuArea        string `json:"asurakuArea"`
			ReviewCount        int    `json:"reviewCount"`
			ShopURL            string `json:"shopUrl"`
			CreditCardFlag     int    `json:"creditCardFlag"`
			ReviewAverage      string `json:"reviewAverage"`
			ShipOverseasArea   string `json:"shipOverseasArea"`
			GenreID            string `json:"genreId"`
			PointRateStartTime string `json:"pointRateStartTime"`
			ItemURL            string `json:"itemUrl"`
		} `json:"Item"`
	} `json:"Items"`
	Title         string `json:"title"`
	LastBuildDate string `json:"lastBuildDate"`
}

func (s *IchibaService) Ranking(ctx context.Context, opt *IchibaRankingParams) (*IchibaRankingResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("IchibaItem/Ranking/20170628")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	ichibaRankingResp := &IchibaRankingResponse{}
	resp, err := s.client.Do(ctx, req, ichibaRankingResp)
	if err != nil {
		return nil, resp, err
	}

	return ichibaRankingResp, resp, nil
}
