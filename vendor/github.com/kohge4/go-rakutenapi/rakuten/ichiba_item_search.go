package rakuten

import (
	"context"
	"fmt"
)

type IchibaItemSearchParams struct {
	Keyword                 string  `url:"keyword,omitempty"`
	ShopCode                string  `url:"shopCode,omitempty"`
	ItemCode                string  `url:"itemCode,omitempty"`
	GenreID                 int     `url:"genreId,omitempty"`
	TagID                   int     `url:"tagId,omitempty"`
	Hits                    int     `url:"hits,omitempty"`
	Page                    int     `url:"page,omitempty"`
	Sort                    string  `url:"sort,omitempty"`
	MinPrice                int     `url:"minPrice,omitempty"`
	MaxPrice                int     `url:"maxPrice,omitempty"`
	Availability            int     `url:"availability,omitempty"`
	Field                   int     `url:"field,omitempty"`
	Carrier                 int     `url:"carrier,omitempty"`
	ImageFlag               int     `url:"imageflag,omitempty"`
	OrFlag                  int     `url:"orFlag,omitempty"`
	NgKeyWord               string  `url:"NGKeyword,omitempty"`
	PurchaseType            int     `url:"purchasetype,omitempty"`
	ShipOverseasFlag        int     `url:"shipoverseasflag,omitempty"`
	ShipOverseasArea        string  `url:"shipoverseasflag,omitempty"`
	AsurakuFlag             int     `url:"asurakuFlag,omitempty"`
	AsurakuArea             int     `url:"asurakuarea,omitempty"`
	PointRateFlag           int     `url:"pointrateflag,omitempty"`
	PointRate               int     `url:"pointrate,omitempty"`
	PostageFlag             int     `url:"pointageflag,omitempty"`
	CreditCardFlag          int     `url:"creditcardflag,omitempty"`
	GiftFlag                int     `url:"giftflag,omitempty"`
	HasRewviewFlag          int     `url:"hasreviewflag,omitempty"`
	MaxAffiliateRate        float64 `url:"maxaffiliaterate,omitempty"`
	MinAffiliateRate        float64 `url:"minaffiliaterate,omitempty"`
	HasMovieFlag            int     `url:"hasmovieflag,omitempty"`
	PampheletFlag           int     `url:"pampheleflag,omitempty"`
	AppointDeliveryDataFlag int     `url:"appointdeliverydataflag,omitempty"`
	GenreInformationFlag    int     `url:"genreInformationFlag,omitempty"`
	TagInformationFlag      int     `url:"tagInformationFlag,omitempty"`
}

// Item is for
type ItemContent struct {
	MediumImageUrls []struct {
		ImageURL string `json:"imageUrl"`
	} `json:"mediumImageUrls"`
	PointRate         int           `json:"pointRate"`
	ShopOfTheYearFlag int           `json:"shopOfTheYearFlag"`
	AffiliateRate     float64       `json:"affiliateRate"`
	ShipOverseasFlag  int           `json:"shipOverseasFlag"`
	AsurakuFlag       int           `json:"asurakuFlag"`
	EndTime           string        `json:"endTime"`
	TaxFlag           int           `json:"taxFlag"`
	StartTime         string        `json:"startTime"`
	ItemCaption       string        `json:"itemCaption"`
	Catchcopy         string        `json:"catchcopy"`
	TagIds            []interface{} `json:"tagIds"`
	SmallImageUrls    []struct {
		ImageURL string `json:"imageUrl"`
	} `json:"smallImageUrls"`
	AsurakuClosingTime string  `json:"asurakuClosingTime"`
	ImageFlag          int     `json:"imageFlag"`
	Availability       int     `json:"availability"`
	ShopAffiliateURL   string  `json:"shopAffiliateUrl"`
	ItemCode           string  `json:"itemCode"`
	PostageFlag        int     `json:"postageFlag"`
	ItemName           string  `json:"itemName"`
	ItemPrice          int     `json:"itemPrice"`
	PointRateEndTime   string  `json:"pointRateEndTime"`
	ShopCode           string  `json:"shopCode"`
	AffiliateURL       string  `json:"affiliateUrl"`
	GiftFlag           int     `json:"giftFlag"`
	ShopName           string  `json:"shopName"`
	ReviewCount        int     `json:"reviewCount"`
	AsurakuArea        string  `json:"asurakuArea"`
	ShopURL            string  `json:"shopUrl"`
	CreditCardFlag     int     `json:"creditCardFlag"`
	ReviewAverage      float64 `json:"reviewAverage"`
	ShipOverseasArea   string  `json:"shipOverseasArea"`
	GenreID            string  `json:"genreId"`
	PointRateStartTime string  `json:"pointRateStartTime"`
	ItemURL            string  `json:"itemUrl"`
}

type Item struct {
	Item ItemContent `json:"Item"`
}

// IchibaItemResponse is response の type 必要かも
type IchibaItemResponse struct {
	Items            []Item
	PageCount        int           `json:"pageCount"`
	TagInformation   []interface{} `json:"TagInformation"`
	Hits             int           `json:"hits"`
	Last             int           `json:"last"`
	Count            int           `json:"count"`
	Page             int           `json:"page"`
	Carrier          int           `json:"carrier"`
	GenreInformation []interface{} `json:"GenreInformation"`
	First            int           `json:"first"`
}

func (s *IchibaService) Search(ctx context.Context, opt *IchibaItemSearchParams) (*IchibaItemResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("IchibaItem/Search/20170706?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	ichibaItemResp := &IchibaItemResponse{}
	resp, err := s.client.Do(ctx, req, ichibaItemResp)
	if err != nil {
		return nil, resp, err
	}
	return ichibaItemResp, resp, nil
}
