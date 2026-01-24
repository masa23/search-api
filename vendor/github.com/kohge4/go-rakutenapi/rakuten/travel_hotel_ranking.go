package rakuten

import (
	"context"
	"fmt"
)

type TravelHotelRankingParams struct {
	// Genre: onsen, all, premium   Carrier:0
	Genre   string `url:"genre,omitempty"`
	Carrier int    `url:"carrier,omitempty"`
}

type TravelHotelRankingResponse struct {
	Rankings []struct {
		Ranking struct {
			Genre         string `json:"genre"`
			Title         string `json:"title"`
			LastBuildDate string `json:"lastBuildDate"`
			Hotels        []struct {
				Hotel struct {
					Rank                int         `json:"rank"`
					HotelNo             int         `json:"hotelNo"`
					HotelName           string      `json:"hotelName"`
					MiddleClassName     string      `json:"middleClassName"`
					UserReview          string      `json:"userReview"`
					ReviewCount         int         `json:"reviewCount"`
					HotelInformationURL string      `json:"hotelInformationUrl"`
					PlanListURL         string      `json:"planListUrl"`
					CheckAvailableURL   string      `json:"checkAvailableUrl"`
					ReviewURL           string      `json:"reviewUrl"`
					HotelImageURL       string      `json:"hotelImageUrl"`
					HotelThumbnailURL   interface{} `json:"hotelThumbnailUrl"`
					ReviewAverage       float64     `json:"reviewAverage"`
					Carrier             int         `json:"carrier"`
				} `json:"hotel"`
			} `json:"hotels"`
		} `json:"Ranking"`
	} `json:"Rankings"`
}

func (s *TravelService) HotelRanking(ctx context.Context, opt *TravelHotelRankingParams) (*TravelHotelRankingResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Travel/HotelRanking/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &TravelHotelRankingResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
