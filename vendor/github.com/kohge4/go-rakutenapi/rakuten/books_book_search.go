package rakuten

import (
	"context"
	"fmt"
)

type BookSearchParams struct {
	Title                string `url:"title,omitempty"`
	Author               string `url:"author,omitempty"`
	PublisherName        string `url:"publisherName,omitempty"`
	Size                 int    `url:"size,omitempty"`
	IsbnJan              string `url:"isbnJan,omitempty"`
	BooksGenreID         string `url:"booksGenreId,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Availability         int    `url:"availability,omitempty"`
	OutOfStockFlag       int    `url:"outOfStockId,omitempty"`
	ChirayomiFlag        string `url:"chirayomiFlag,omitempty"`
	Sort                 string `url:"sort,omitempty"`
	LimitedFlag          int    `url:"lomitedFlag,omitempty"`
	Field                int    `url:"filed,omitempty"`
	Carrier              int    `url:"carrier,omitempty"`
	GenreInformationFlag int    `url:"genreInformationFlag,omitempty"`
}

type BookSearchResponse struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	Carrier   int `json:"carrier"`
	PageCount int `json:"pageCount"`
	Items     []struct {
		Item struct {
			Title          string `json:"title"`
			TitleKana      string `json:"titleKana"`
			SubTitle       string `json:"subTitle"`
			SubTitleKana   string `json:"subTitleKana"`
			SeriesName     string `json:"seriesName"`
			SeriesNameKana string `json:"seriesNameKana"`
			Contents       string `json:"contents"`
			Author         string `json:"author"`
			AuthorKana     string `json:"authorKana"`
			PublisherName  string `json:"publisherName"`
			Size           string `json:"size"`
			Isbn           string `json:"isbn"`
			ItemCaption    string `json:"itemCaption"`
			SalesDate      string `json:"salesDate"`
			ItemPrice      int    `json:"itemPrice"`
			ListPrice      int    `json:"listPrice"`
			DiscountRate   int    `json:"discountRate"`
			DiscountPrice  int    `json:"discountPrice"`
			ItemURL        string `json:"itemUrl"`
			AffiliateURL   string `json:"affiliateUrl"`
			SmallImageURL  string `json:"smallImageUrl"`
			MediumImageURL string `json:"mediumImageUrl"`
			LargeImageURL  string `json:"largeImageUrl"`
			ChirayomiURL   string `json:"chirayomiUrl"`
			Availability   string `json:"availability"`
			PostageFlag    int    `json:"postageFlag"`
			LimitedFlag    int    `json:"limitedFlag"`
			ReviewCount    int    `json:"reviewCount"`
			ReviewAverage  string `json:"reviewAverage"`
			BooksGenreID   string `json:"booksGenreId"`
		} `json:"Item"`
	} `json:"Items"`
	GenreInformation []interface{} `json:"GenreInformation"`
}

func (s *BooksService) BookSearch(ctx context.Context, opt *BookSearchParams) (*BookSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("BooksBook/Search/20170404?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &BookSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
