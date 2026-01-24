package rakuten

import (
	"context"
	"fmt"
)

type FavoDeleteResponse struct {
}

type FavoDeleteParams struct {
	AccessToken string `url:"access_token,omitempty"`
	Format      string `url:"format,omitempty"`
	ItemCode    string `url:"itemCode,optempty"`
}

func (s *FavoService) Delete(ctx context.Context, opt *FavoDeleteParams) (*FavoDeleteResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("FavoriteBookmark/Delete/20120627")
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

	favoDeleteResp := &FavoDeleteResponse{}
	resp, err := s.client.Do(ctx, req, favoDeleteResp)
	if err != nil {
		return nil, resp, err
	}
	return favoDeleteResp, resp, nil
}
