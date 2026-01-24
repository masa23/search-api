package rakuten

import (
	"context"
	"fmt"
)

type TravelKeywordHotelSearchParams struct {
	Carrier             int    `url:"carrier,omitempty"`
	Page                int    `url:"page,omitempty"`
	Hits                int    `url:"hits,omitempty"`
	DatumType           int    `url:"datumType,omitempty"`
	Keyword             string `url:"keyword,omitempty"`
	MiddleClassCode     string `url:"middleClassCode,omitempty"`
	SearchField         int    `url:"searchField,omitempty"`
	HotelChainCode      string `url:"hotelChainCode,omitempty"`
	HotelcThumbnailSize int    `url:"hotelThumbnailSize,omitempty"`
	ResponseType        string `url:"responseType,omitempty"`
	Sort                string `url:"sort,omitempty"`
}

type TravelKeywordHotelSearchResponse struct {
	PagingInfo struct {
		RecordCount int `json:"recordCount"`
		PageCount   int `json:"pageCount"`
		Page        int `json:"page"`
		First       int `json:"first"`
		Last        int `json:"last"`
	} `json:"pagingInfo"`
	Hotels []struct {
		Hotel []struct {
			HotelBasicInfo struct {
				HotelNo             int         `json:"hotelNo"`
				HotelName           string      `json:"hotelName"`
				HotelInformationURL string      `json:"hotelInformationUrl"`
				PlanListURL         string      `json:"planListUrl"`
				DpPlanListURL       string      `json:"dpPlanListUrl"`
				ReviewURL           string      `json:"reviewUrl"`
				HotelKanaName       string      `json:"hotelKanaName"`
				HotelSpecial        string      `json:"hotelSpecial"`
				HotelMinCharge      int         `json:"hotelMinCharge"`
				Latitude            float64     `json:"latitude"`
				Longitude           float64     `json:"longitude"`
				PostalCode          string      `json:"postalCode"`
				Address1            string      `json:"address1"`
				Address2            string      `json:"address2"`
				TelephoneNo         string      `json:"telephoneNo"`
				FaxNo               string      `json:"faxNo"`
				Access              string      `json:"access"`
				ParkingInformation  string      `json:"parkingInformation"`
				NearestStation      string      `json:"nearestStation"`
				HotelImageURL       string      `json:"hotelImageUrl"`
				HotelThumbnailURL   string      `json:"hotelThumbnailUrl"`
				RoomImageURL        interface{} `json:"roomImageUrl"`
				RoomThumbnailURL    interface{} `json:"roomThumbnailUrl"`
				HotelMapImageURL    string      `json:"hotelMapImageUrl"`
				ReviewCount         int         `json:"reviewCount"`
				ReviewAverage       float64     `json:"reviewAverage"`
				UserReview          string      `json:"userReview"`
			} `json:"hotelBasicInfo,omitempty"`
			HotelRatingInfo struct {
				ServiceAverage   float64 `json:"serviceAverage"`
				LocationAverage  float64 `json:"locationAverage"`
				RoomAverage      float64 `json:"roomAverage"`
				EquipmentAverage float64 `json:"equipmentAverage"`
				BathAverage      float64 `json:"bathAverage"`
				MealAverage      float64 `json:"mealAverage"`
			} `json:"hotelRatingInfo,omitempty"`
		} `json:"hotel"`
	} `json:"hotels"`
}

func (s *TravelService) KeywordHotelSearch(ctx context.Context, opt *TravelKeywordHotelSearchParams) (*TravelKeywordHotelSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Travel/KeywordHotelSearch/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &TravelKeywordHotelSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
