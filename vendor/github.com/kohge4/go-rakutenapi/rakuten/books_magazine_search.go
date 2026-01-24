package rakuten

import (
	"context"
	"fmt"
)

type BooksMagazineSearchParams struct {
	Title                string `url:"title,omitempty"`
	Publishername        string `url:"publisherName,omitempty"`
	BooksGenreID         string `url:"booksGenreID,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Availability         int    `url:"availability,omitempty"`
	OutOfStockFlag       int    `url:"outOfStock,omitempty"`
	ChirayomiFlag        string `url:"chirayomiFlag,omitempty"`
	Sort                 string `url:"sort,omitempty"`
	LimitedFlag          int    `url:"limitedFlag,omitempty"`
	Field                int    `url:"filed,omitempty"`
	Carrier              int    `url:"carrier,omitempty"`
	GenreInformationFlag int    `url:"genreInfoamtionFlag,omitempty"`
}

type BooksMagazineSearchResponse struct {
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
			PublisherName  string `json:"publisherName"`
			Jan            string `json:"jan"`
			ItemCaption    string `json:"itemCaption"`
			SalesDate      string `json:"salesDate"`
			Cycle          string `json:"cycle"`
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

func (s *BooksService) MagazineSearch(ctx context.Context, opt *BooksMagazineSearchParams) (*BooksMagazineSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("BooksMagazine/Search/20170404?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &BooksMagazineSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
