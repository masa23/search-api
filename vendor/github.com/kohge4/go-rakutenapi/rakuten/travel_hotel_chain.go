package rakuten

import (
	"context"
	"fmt"
)

type TravelHotelChainParams struct{}

type TravelHotelChainResponse struct {
	LargeClasses []struct {
		LargeClass []struct {
			LargeClassCode string `json:"largeClassCode"`
			HotelChains    []struct {
				HotelChain struct {
					HotelChainCode     string `json:"hotelChainCode"`
					HotelChainName     string `json:"hotelChainName"`
					HotelChainNameKana string `json:"hotelChainNameKana"`
					HotelChainComment  string `json:"hotelChainComment"`
				} `json:"hotelChain"`
			} `json:"hotelChains"`
		} `json:"largeClass"`
	} `json:"largeClasses"`
}

func (s *TravelService) HotelChain(ctx context.Context, opt *TravelHotelChainParams) (*TravelHotelChainResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Travel/GetHotelChainList/20131024?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &TravelHotelChainResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
