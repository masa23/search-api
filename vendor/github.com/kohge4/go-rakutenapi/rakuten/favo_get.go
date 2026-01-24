package rakuten

import (
	"context"
	"fmt"
)

type FavoGetResponse struct {
	Summary struct {
		Count     int `json:"count"`
		Hits      int `json:"hits"`
		PageCount int `json:"pageCount"`
	} `json:"summary"`
	Items []struct {
		Item struct {
			BookmarkID     string `json:"bookmarkId"`
			ItemCode       string `json:"itemCode"`
			ProductID      string `json:"productId"`
			ShopName       string `json:"shopName"`
			ShopURL        string `json:"shopUrl"`
			ItemName       string `json:"itemName"`
			ItemURL        string `json:"itemUrl"`
			SmallImageURL  string `json:"smallImageUrl"`
			MediumImageURL string `json:"mediumImageUrl"`
			ReviewCount    int    `json:"reviewCount"`
			ReviewURL      string `json:"reviewUrl"`
			PointRate      int    `json:"pointRate"`
			ReviewAverage  string `json:"reviewAverage"`
			PostageFlag    int    `json:"postageFlag"`
			TaxFlag        int    `json:"taxFlag"`
			AffiliateURL   string `json:"affiliateUrl"`
		} `json:"item"`
	} `json:"items"`
}

type FavoGetParams struct {
	AccessToken string `url:"access_token,omitempty"`
	Format      string `url:"format,omitempty"`
}

func (s *FavoService) Get(ctx context.Context, opt *FavoGetParams) (*FavoGetResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("FavoriteBookmark/List/20170426")
	param, err := addOptions(urlSuffix, opt)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("\nparam\n")
	fmt.Println(param)
	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}

	favoGetResp := &FavoGetResponse{}
	resp, err := s.client.Do(ctx, req, favoGetResp)
	if err != nil {
		return nil, resp, err
	}
	return favoGetResp, resp, nil
}
