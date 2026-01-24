package rakuten

import (
	"context"
	"fmt"
)

type TravelVacantHotelSearchParams struct {
	LargeClassCode      string  `url:"largeClassCode,omitempty"`
	MiddleClassCode     string  `url:"middleClassCode,omitempty"`
	SmallClassCode      string  `url:"smallClassCode,omitempty"`
	DetailClassCode     string  `url:"detailClassCode,omitempty"`
	HotelNo             int     `url:"hotelNo,omitempty"`
	CheckDate           string  `url:"checkDate,omitempty"`
	CheckOutDate        string  `url:"checkOutDate,omitempty"`
	AdultNum            int     `url:"adultNum,omitempty"`
	UpClassNum          int     `url:"upClassNum,omitempty"`
	LowClassNum         int     `url:"lowClassNum,omitempty"`
	InfrantWithMBNum    int     `url:"infrantWithMBNum,omitempty"`
	InfrantWithMNum     int     `url:"infrantWithMNum,omitempty"`
	InfrantWithBNum     int     `url:"infrantWithBNum,omitempty"`
	InfrantWithoutMBNum int     `url:"infrantWithoutMBNum,omitempty"`
	RoomNum             int     `url:"roomNum,omitempty"`
	MaxCharge           int     `url:"maxCharge,omitempty"`
	Latitue             float64 `url:"latitude,omitempty"`
	Longtitude          float64 `url:"longtitude,omitempty"`
	SearchRadius        int     `url:"searchRadius,omitempty"`
	SqueezeCondition    string  `url:"squeezeCondition,omitempty"`
	Carrier             int     `url:"carrier,omitempty"`
	DatumType           int     `url:"datumType,omitempty"`
	Hits                int     `url:"hits,omitempty"`
	Page                int     `url:"page,omitempty"`
	SearchPattern       int     `url:"searchPattern,omitempty"`
	HotelThumbnailSize  int     `url:"hotelThumbnailSize,omitempty"`
	ResponseType        string  `url:"responseType,omitempty"`
	Sort                string  `url:"sort,omitempty"`
	AllReturnFlag       int     `url:"allReturnFlag,omitempty"`
}

type TravelVacantHotelSearchResponse struct {
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
				HotelNo             int     `json:"hotelNo"`
				HotelName           string  `json:"hotelName"`
				HotelInformationURL string  `json:"hotelInformationUrl"`
				PlanListURL         string  `json:"planListUrl"`
				DpPlanListURL       string  `json:"dpPlanListUrl"`
				ReviewURL           string  `json:"reviewUrl"`
				HotelKanaName       string  `json:"hotelKanaName"`
				HotelSpecial        string  `json:"hotelSpecial"`
				HotelMinCharge      int     `json:"hotelMinCharge"`
				Latitude            float64 `json:"latitude"`
				Longitude           float64 `json:"longitude"`
				PostalCode          string  `json:"postalCode"`
				Address1            string  `json:"address1"`
				Address2            string  `json:"address2"`
				TelephoneNo         string  `json:"telephoneNo"`
				FaxNo               string  `json:"faxNo"`
				Access              string  `json:"access"`
				ParkingInformation  string  `json:"parkingInformation"`
				NearestStation      string  `json:"nearestStation"`
				HotelImageURL       string  `json:"hotelImageUrl"`
				HotelThumbnailURL   string  `json:"hotelThumbnailUrl"`
				RoomImageURL        string  `json:"roomImageUrl"`
				RoomThumbnailURL    string  `json:"roomThumbnailUrl"`
				HotelMapImageURL    string  `json:"hotelMapImageUrl"`
				ReviewCount         int     `json:"reviewCount"`
				ReviewAverage       float64 `json:"reviewAverage"`
				UserReview          string  `json:"userReview"`
			} `json:"hotelBasicInfo,omitempty"`
			RoomInfo []struct {
				RoomBasicInfo struct {
					RoomClass           string `json:"roomClass"`
					RoomName            string `json:"roomName"`
					PlanID              int    `json:"planId"`
					PlanName            string `json:"planName"`
					PointRate           int    `json:"pointRate"`
					WithDinnerFlag      int    `json:"withDinnerFlag"`
					DinnerSelectFlag    int    `json:"dinnerSelectFlag"`
					WithBreakfastFlag   int    `json:"withBreakfastFlag"`
					BreakfastSelectFlag int    `json:"breakfastSelectFlag"`
					Payment             string `json:"payment"`
					ReserveURL          string `json:"reserveUrl"`
					SalesformFlag       int    `json:"salesformFlag"`
				} `json:"roomBasicInfo,omitempty"`
				DailyCharge struct {
					StayDate      string `json:"stayDate"`
					RakutenCharge int    `json:"rakutenCharge"`
					Total         int    `json:"total"`
					ChargeFlag    int    `json:"chargeFlag"`
				} `json:"dailyCharge,omitempty"`
			} `json:"roomInfo,omitempty"`
		} `json:"hotel"`
	} `json:"hotels"`
}

func (s *TravelService) VacantHotelSearch(ctx context.Context, opt *TravelVacantHotelSearchParams) (*TravelVacantHotelSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Travel/VacantHotelSearch/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &TravelVacantHotelSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
