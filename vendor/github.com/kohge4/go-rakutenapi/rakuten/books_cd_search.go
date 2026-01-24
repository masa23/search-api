package rakuten

import (
	"context"
	"fmt"
)

type BooksCDSearchParams struct {
	Title                string `url:"title,omitempty"`
	ArtistName           string `url:"artistName,omitempty"`
	Label                string `url:"label,omitempty"`
	Jan                  int    `url:"jan,omitempty"`
	BooksGenreID         string `url:"booksGenreId,omitempty"`
	Hits                 int    `url:"hits,omitempty"`
	Page                 int    `url:"page,omitempty"`
	Availability         int    `url:"availability,omitempty"`
	OutOfStockFlag       int    `url:"outOfStockFlag,omitempty"`
	Sort                 string `url:"sort,omitempty"`
	LimitedFlag          int    `url:"limitedFlag,omitempty"`
	Carrier              int    `url:"carrier,omitempty"`
	GenreInformationFlag int    `url:"genreInformationFlag,omitempty"`
}

type BooksCDSearchResponse struct {
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
			ArtistName     string `json:"artistName"`
			ArtistNameKana string `json:"artistNameKana"`
			Label          string `json:"label"`
			Size           string `json:"size"`
			Jan            string `json:"jan"`
			MakerCode      string `json:"makerCode"`
			ItemCaption    string `json:"itemCaption"`
			PlayList       string `json:"playList"`
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

func (s *BooksService) CDSearch(ctx context.Context, opt *BooksCDSearchParams) (*BooksCDSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("BooksCD/Search/20170404?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &BooksCDSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
