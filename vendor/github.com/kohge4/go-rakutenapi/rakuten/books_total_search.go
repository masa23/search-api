package rakuten

import (
	"context"
	"fmt"
)

// 仕様: https://webservice.rakuten.co.jp/api/bookstotalsearch/

type BooksTotalSearchParams struct {
	Keyword              string `url:"keyword,omitempty"`
	BooksGenreID         string `url:"booksGenreId,omitempty"`
	IsbnJan              string `url:"isbnJan,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Sort                 string `url:"sort,omitempty"`
	Availability         int    `url:"availability,omitempty"`
	OutOfStockFlag       int    `url:"outOfStockFlag,omitempty"`
	ChirayomiFlag        string `url:"chirayomiFlag,omitempty"`
	LimitedFlag          int    `url:"limitedFlag,omitempty"`
	Field                int    `url:"filed,omitempty"`
	Carrier              int    `url:"carrier,omitempty"`
	OrFlag               int    `url:"orFlag,omitempty"`
	NGKeyword            string `url:"NGKeyword,omitempty"`
	GenreInformationFlag int    `url:"genreInformationFlag,omitempty"`
}

type BooksTotalSearchResponse struct {
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
			Author         string `json:"author"`
			ArtistName     string `json:"artistName"`
			PublisherName  string `json:"publisherName"`
			Label          string `json:"label"`
			Isbn           string `json:"isbn"`
			Jan            string `json:"jan"`
			Hardware       string `json:"hardware"`
			Os             string `json:"os"`
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

func (s *BooksService) TotalSearch(ctx context.Context, opt *BooksTotalSearchParams) (*BooksTotalSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("BooksTotal/Search/20170404?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &BooksTotalSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
