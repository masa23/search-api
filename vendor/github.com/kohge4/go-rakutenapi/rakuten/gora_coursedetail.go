package rakuten

import (
	"context"
	"fmt"
)

type GolfCourseDetailParams struct {
	GolfCourseID int `url:"golfCourseId,omitempty"`
	Carrier      int `url:"carrier,omitempty"`
}

type GolfCourseDetailResponse struct {
	Item struct {
		Carrier                int     `json:"carrier"`
		GolfCourseID           int     `json:"golfCourseId"`
		GolfCourseName         string  `json:"golfCourseName"`
		GolfCourseAbbr         string  `json:"golfCourseAbbr"`
		GolfCourseNameKana     string  `json:"golfCourseNameKana"`
		GolfCourseCaption      string  `json:"golfCourseCaption"`
		Information            string  `json:"information"`
		Highway                string  `json:"highway"`
		Ic                     string  `json:"ic"`
		IcDistance             string  `json:"icDistance"`
		Latitude               string  `json:"latitude"`
		Longitude              string  `json:"longitude"`
		PostalCode             string  `json:"postalCode"`
		Address                string  `json:"address"`
		TelephoneNo            string  `json:"telephoneNo"`
		FaxNo                  string  `json:"faxNo"`
		OpenDay                string  `json:"openDay"`
		CloseDay               string  `json:"closeDay"`
		CreditCard             string  `json:"creditCard"`
		Shoes                  string  `json:"shoes"`
		DressCode              string  `json:"dressCode"`
		PracticeFacility       string  `json:"practiceFacility"`
		LodgingFacility        string  `json:"lodgingFacility"`
		OtherFacility          string  `json:"otherFacility"`
		GolfCourseImageURL1    string  `json:"golfCourseImageUrl1"`
		GolfCourseImageURL2    string  `json:"golfCourseImageUrl2"`
		GolfCourseImageURL3    string  `json:"golfCourseImageUrl3"`
		GolfCourseImageURL4    string  `json:"golfCourseImageUrl4"`
		GolfCourseImageURL5    string  `json:"golfCourseImageUrl5"`
		WeekdayMinPrice        int     `json:"weekdayMinPrice"`
		BaseWeekdayMinPrice    int     `json:"baseWeekdayMinPrice"`
		HolidayMinPrice        int     `json:"holidayMinPrice"`
		BaseHolidayMinPrice    int     `json:"baseHolidayMinPrice"`
		Designer               string  `json:"designer"`
		CourseType             string  `json:"courseType"`
		CourseVerticalInterval string  `json:"courseVerticalInterval"`
		Dimension              string  `json:"dimension"`
		Green                  string  `json:"green"`
		GreenCount             string  `json:"greenCount"`
		HoleCount              int     `json:"holeCount"`
		ParCount               int     `json:"parCount"`
		CourseName             string  `json:"courseName"`
		CourseDistance         string  `json:"courseDistance"`
		LongDrivingContest     string  `json:"longDrivingContest"`
		NearPin                string  `json:"nearPin"`
		RatingNum              int     `json:"ratingNum"`
		Evaluation             float64 `json:"evaluation"`
		Staff                  float64 `json:"staff"`
		Facility               float64 `json:"facility"`
		Meal                   float64 `json:"meal"`
		Course                 float64 `json:"course"`
		Costperformance        float64 `json:"costperformance"`
		Distance               float64 `json:"distance"`
		Fairway                float64 `json:"fairway"`
		ReserveCalURL          string  `json:"reserveCalUrl"`
		VoiceURL               string  `json:"voiceUrl"`
		LayoutURL              string  `json:"layoutUrl"`
		RouteMapURL            string  `json:"routeMapUrl"`
		NewPlans               []struct {
			Plan struct {
				Month        string `json:"month"`
				PlanName     string `json:"planName"`
				PlanDate     string `json:"planDate"`
				Service      string `json:"service"`
				Price        string `json:"price"`
				BasePrice    int    `json:"basePrice"`
				SalesTax     int    `json:"salesTax"`
				CourseUseTax int    `json:"courseUseTax"`
				OtherTax     int    `json:"otherTax"`
			} `json:"plan"`
		} `json:"newPlans"`
		Ratings []struct {
			Rating struct {
				Title           string `json:"title"`
				NickName        string `json:"nickName"`
				Prefecture      string `json:"prefecture"`
				Age             string `json:"age"`
				Sex             string `json:"sex"`
				Times           int    `json:"times"`
				Evaluation      int    `json:"evaluation"`
				Staff           int    `json:"staff"`
				Facility        int    `json:"facility"`
				Meal            int    `json:"meal"`
				Course          int    `json:"course"`
				Costperformance int    `json:"costperformance"`
				Distance        int    `json:"distance"`
				Fairway         int    `json:"fairway"`
				Comment         string `json:"comment"`
			} `json:"rating"`
		} `json:"ratings"`
	} `json:"Item"`
}

func (s GoraService) GolfCourseDetailSearch(ctx context.Context, opt *GolfCourseDetailParams) (*GolfCourseDetailResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Gora/GoraGolfCourseDetail/20170623?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	goraResp := &GolfCourseDetailResponse{}
	resp, err := s.client.Do(ctx, req, goraResp)
	if err != nil {
		return nil, resp, err
	}

	return goraResp, resp, nil
}
