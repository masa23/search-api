package rakuten

import (
	"context"
	"fmt"
)

type GolfPlanParams struct {
	GolfCourseName   string `url:"golfCourseName,omitempty"`
	AreaCode         int    `url:"areaCode,omitempty"`
	GolfCourseID     int    `url:"golfCourseId,omitempty"`
	PlayDate         string `url:"playDate,omitempty"`
	Hits             int    `url:"hits,omitempty"`
	Page             int    `url:"page,omitempty"`
	Sort             string `url:"sort,omitempty"`
	MinPrice         int    `url:"minPrice,omitempty"`
	MaxPrice         int    `url:"maxPrice,omitempty"`
	StartTimeZone    int    `url:"startTimeZone,omitempty"`
	Range            int    `url:"range,omitempty"`
	PalnCaddie       int    `url:"planCaddie,omitempty"`
	PlanCart         int    `url:"planCart,omitempty"`
	PlanStay         int    `url:"planStay,omitempty"`
	PlanLunch        int    `url:"planLanch,omitempty"`
	Plan2Sum         int    `url:"plam2sum,omitempty"`
	PlanDiscount4Sum int    `url:"planDiscount4sum,omitempty"`
	PlanAdd1RFree    int    `url:"planAdd1RFree,omitempty"`
	PlanAddHalfFree  int    `url:"planAddHalfFree,omitempty"`
	PlanNoAddDee2b   int    `url:"planNoAddFee2b,omitempty"`
	PlanNoAddDee3b   int    `url:"planNoAddFee3b,omitempty"`
	PlanDrink        int    `url:"planDrink,omitempty"`
	PlanGoraOrg      int    `url:"planGoraOrg,omitempty"`
	PlanLesson       int    `url:"planLesson,omitempty"`
	PlanOpenCompe    int    `url:"planOpenCompe,omitempty"`
	PlanHalfRound    int    `url:"planHalfRound,omitempty"`
	PlanEarly        int    `url:"planEarly,omitempty"`
	NGPlan           string `url:"NGPlan,omitempty"`
	Coursetype       int    `url:"courseType,omitempty"`
	ShapeWideFairway int    `url:"shapeWideFairway,omitempty"`
	ShapeLessOB      int    `url:"shapeLessOB,omitempty"`
	PracticeFacility int    `url:"practiceFacility,omitempty"`
	LongingFacility  int    `url:"longingFacility,omitempty"`
	OpenFacility     int    `url:"openFacility,omitempty"`
	HighwayCode      int    `url:"highwayCode,omitempty"`
	IcDistance       int    `url:"icDistance,omitempty"`
	PointFlag        int    `url:"pointFlag,omitempty"`
}

type GolfPlanResponse struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	PageCount int `json:"pageCount"`
	Items     []struct {
		Item struct {
			GolfCourseID               int         `json:"golfCourseId"`
			GolfCourseName             string      `json:"golfCourseName"`
			GolfCourseCaption          string      `json:"golfCourseCaption"`
			GolfCourseRsvType          int         `json:"golfCourseRsvType"`
			AreaCode                   int         `json:"areaCode"`
			Prefecture                 string      `json:"prefecture"`
			HighwayCode                int         `json:"highwayCode"`
			Highway                    string      `json:"highway"`
			Ic                         string      `json:"ic"`
			IcDistance                 string      `json:"icDistance"`
			GolfCourseImageURL         string      `json:"golfCourseImageUrl"`
			DisplayWeekdayMinPrice     interface{} `json:"displayWeekdayMinPrice"`
			DisplayWeekdayMinBasePrice interface{} `json:"displayWeekdayMinBasePrice"`
			DisplayHolidayMinPrice     string      `json:"displayHolidayMinPrice"`
			DisplayHolidayMinBasePrice string      `json:"displayHolidayMinBasePrice"`
			CancelFeeFlag              int         `json:"cancelFeeFlag"`
			CancelFee                  string      `json:"cancelFee"`
			RatingNum                  int         `json:"ratingNum"`
			Evaluation                 float64     `json:"evaluation"`
			ReserveCalURLPC            string      `json:"reserveCalUrlPC"`
			ReserveCalURLMobile        string      `json:"reserveCalUrlMobile"`
			RatingURLPC                string      `json:"ratingUrlPC"`
			RatingURLMobile            string      `json:"ratingUrlMobile"`
			PlanInfo                   []struct {
				Plan struct {
					PlanID             int         `json:"planId"`
					PlanName           string      `json:"planName"`
					PlanType           int         `json:"planType"`
					LimitedTimeFlag    int         `json:"limitedTimeFlag"`
					Price              int         `json:"price"`
					BasePrice          int         `json:"basePrice"`
					SalesTax           int         `json:"salesTax"`
					CourseUseTax       int         `json:"courseUseTax"`
					OtherTax           int         `json:"otherTax"`
					PlayerNumMin       int         `json:"playerNumMin"`
					PlayerNumMax       int         `json:"playerNumMax"`
					StartTimeZone      string      `json:"startTimeZone"`
					Round              string      `json:"round"`
					Caddie             int         `json:"caddie"`
					Cart               int         `json:"cart"`
					Assu2Sum           int         `json:"assu2sum"`
					AddFee2BFlag       int         `json:"addFee2bFlag"`
					AddFee2B           int         `json:"addFee2b"`
					Assortment2BFlag   int         `json:"assortment2bFlag"`
					AddFee3BFlag       int         `json:"addFee3bFlag"`
					AddFee3B           int         `json:"addFee3b"`
					Assortment3BFlag   int         `json:"assortment3bFlag"`
					Discount4SumFlag   int         `json:"discount4sumFlag"`
					Lunch              int         `json:"lunch"`
					Drink              int         `json:"drink"`
					Stay               int         `json:"stay"`
					Lesson             int         `json:"lesson"`
					OpenCompe          int         `json:"openCompe"`
					RegularCompe       int         `json:"regularCompe"`
					CompePlayGroupMin  int         `json:"compePlayGroupMin"`
					CompePlayMemberMin int         `json:"compePlayMemberMin"`
					CompePrivilegeFree interface{} `json:"compePrivilegeFree"`
					CompeOption        interface{} `json:"compeOption"`
					Other              string      `json:"other"`
					PointFlag          int         `json:"pointFlag"`
					Point              int         `json:"point"`
					CallInfo           struct {
						PlayDate             string `json:"playDate"`
						StockStatus          int    `json:"stockStatus"`
						StockCount           int    `json:"stockCount"`
						ReservePageURLPC     string `json:"reservePageUrlPC"`
						ReservePageURLMobile string `json:"reservePageUrlMobile"`
					} `json:"callInfo"`
				} `json:"plan"`
			} `json:"planInfo"`
		} `json:"Item"`
	} `json:"Items"`
}

func (s GoraService) GolfPlanSearch(ctx context.Context, opt *GolfPlanParams) (*GolfPlanResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Gora/GoraPlanSearch/20170623?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	goraResp := &GolfPlanResponse{}
	resp, err := s.client.Do(ctx, req, goraResp)
	if err != nil {
		return nil, resp, err
	}

	return goraResp, resp, nil
}
