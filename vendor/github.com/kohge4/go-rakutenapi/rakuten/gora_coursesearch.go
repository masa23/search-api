package rakuten

import (
	"context"
	"fmt"
)

type GolfCourseParams struct {
	Keyword      string `url:"keyword,omitempty"`
	AreaCode     int    `url:"areaCode,omitempty"`
	Latitude     int    `url:"latitude,omitempty"`
	Longtitude   int    `url:"longtitude,omitempty"`
	SearchRadius int    `url:"searchRadius,omitempty"`
	Hits         int    `url:"hits,omitempty"`
	Page         int    `url:"page,omitempty"`
	Sort         string `url:"sort,omitempty"`
	Reservation  int    `url:"reservation,omitempty"`
	Carrier      int    `url:"carrier,omitempty"`
}

type GolfCourseResponse struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	Carrier   int `json:"carrier"`
	PageCount int `json:"pageCount"`
	Items     []struct {
		Item struct {
			GolfCourseID        int     `json:"golfCourseId"`
			GolfCourseName      string  `json:"golfCourseName"`
			GolfCourseAbbr      string  `json:"golfCourseAbbr"`
			GolfCourseNameKana  string  `json:"golfCourseNameKana"`
			GolfCourseCaption   string  `json:"golfCourseCaption"`
			Address             string  `json:"address"`
			Latitude            float64 `json:"latitude"`
			Longitude           float64 `json:"longitude"`
			Highway             string  `json:"highway"`
			GolfCourseDetailURL string  `json:"golfCourseDetailUrl"`
			ReserveCalURL       string  `json:"reserveCalUrl"`
			RatingURL           string  `json:"ratingUrl"`
			GolfCourseImageURL  string  `json:"golfCourseImageUrl"`
			Evaluation          float64 `json:"evaluation"`
		} `json:"Item"`
	} `json:"Items"`
}

func (s GoraService) GolfCourseSearch(ctx context.Context, opt *GolfCourseParams) (*GolfCourseResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Gora/GoraGolfCourseSearch/20170623?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	goraResp := &GolfCourseResponse{}
	resp, err := s.client.Do(ctx, req, goraResp)
	if err != nil {
		return nil, resp, err
	}

	return goraResp, resp, nil
}
