package rakuten

import (
	"context"
	"fmt"
)

type KoboEbooksParams struct {
	Keyword              string `url:"keyword,omitempty"`
	Title                string `url:"title,omitempty"`
	Author               string `url:"author,omitempty"`
	PublisherName        string `url:"publisherName,omitempty"`
	KoboGenreID          string `url:"koboGenreId,omitempty"`
	Language             string `url:"language,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Sort                 string `url:"sort,omitempty"`
	NGKeyword            string `url:"NGKeyword,omitempty"`
	Field                int    `url:"field,omitempty"`
	OrFlag               int    `url:"orFlag,omitempty"`
	GenreInformationFlag int    `url:"genreInformationFlag,omitempty"`
	SalesType            int    `url:"salesType,omitempty"`
}

type KoboEbooksResponse struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	PageCount int `json:"pageCount"`
	Items     []struct {
		Item struct {
			Title          string `json:"title"`
			TitleKana      string `json:"titleKana"`
			SubTitle       string `json:"subTitle"`
			SeriesName     string `json:"seriesName"`
			Author         string `json:"author"`
			AuthorKana     string `json:"authorKana"`
			PublisherName  string `json:"publisherName"`
			Language       string `json:"language"`
			SalesDate      string `json:"salesDate"`
			ItemNumber     string `json:"itemNumber"`
			KoboGenreID    string `json:"koboGenreId"`
			ItemCaption    string `json:"itemCaption"`
			ItemPrice      int    `json:"itemPrice"`
			ItemURL        string `json:"itemUrl"`
			AffiliateURL   string `json:"affiliateUrl"`
			SmallImageURL  string `json:"smallImageUrl"`
			MediumImageURL string `json:"mediumImageUrl"`
			LargeImageURL  string `json:"largeImageUrl"`
			ReviewCount    int    `json:"reviewCount"`
			ReviewAverage  string `json:"reviewAverage"`
			SalesType      int    `json:"salesType"`
		} `json:"Item"`
	} `json:"Items"`
	GenreInformation []interface{} `json:"GenreInformation"`
}

func (s KoboService) EbooksSearch(ctx context.Context, opt *KoboEbooksParams) (*KoboEbooksResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Kobo/EbookSearch/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	koboEbooksResp := &KoboEbooksResponse{}
	resp, err := s.client.Do(ctx, req, koboEbooksResp)
	if err != nil {
		return nil, resp, err
	}

	return koboEbooksResp, resp, nil
}
