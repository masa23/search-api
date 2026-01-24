package rakuten

import (
	"context"
	"fmt"
)

type TravelGetAreaParams struct {
}

type TravelGetAreaResponse struct {
	AreaClasses struct {
		LargeClasses []struct {
			LargeClass []struct {
				LargeClassCode string `json:"largeClassCode,omitempty"`
				LargeClassName string `json:"largeClassName,omitempty"`
				MiddleClasses  []struct {
					MiddleClass []struct {
						MiddleClassCode string `json:"middleClassCode,omitempty"`
						MiddleClassName string `json:"middleClassName,omitempty"`
						SmallClasses    []struct {
							SmallClass []struct {
								SmallClassCode string `json:"smallClassCode,omitempty"`
								SmallClassName string `json:"smallClassName,omitempty"`
								DetailClasses  []struct {
									DetailClass struct {
										DetailClassCode string `json:"detailClassCode"`
										DetailClassName string `json:"detailClassName"`
									} `json:"detailClass"`
								} `json:"detailClasses,omitempty"`
							} `json:"smallClass"`
						} `json:"smallClasses,omitempty"`
					} `json:"middleClass"`
				} `json:"middleClasses,omitempty"`
			} `json:"largeClass"`
		} `json:"largeClasses"`
	} `json:"areaClasses"`
}

func (s *TravelService) GetArea(ctx context.Context, opt *TravelGetAreaParams) (*TravelGetAreaResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Travel/GetAreaClass/20131024?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &TravelGetAreaResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
