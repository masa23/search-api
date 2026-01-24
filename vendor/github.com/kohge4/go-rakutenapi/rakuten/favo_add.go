package rakuten

import (
	"context"
	"fmt"
)

// https://app.rakuten.co.jp/services/api/FavoriteBookmark/List/20170426?access_token=IgABNDsSRaZftliujurut6Mpy4J7WUbEEfT1L4UefQMf9l0zaXp&format=json&affiliateId=1a11cd3a.d53d02bd.1a11cd3b.54ce7249&applicationId=1058291654304154113

//access_token=IgABNDsSRaZftliujurut6Mpy4J7WUbEEfT1L4UefQMf9l0zaXp

type FavoAddResponse struct {
}

type FavoAddParams struct {
	AccessToken string `url:"access_token,omitempty"`
	Format      string `url:"format,omitempty"`
	ItemCode    string `url:"itemCode,optempty"`
}

func (s *FavoService) Add(ctx context.Context, opt *FavoAddParams) (*FavoAddResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("FavoriteBookmark/Add/20120627")
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

	favoAddResp := &FavoAddResponse{}
	resp, err := s.client.Do(ctx, req, favoAddResp)
	if err != nil {
		return nil, resp, err
	}
	return favoAddResp, resp, nil
}
