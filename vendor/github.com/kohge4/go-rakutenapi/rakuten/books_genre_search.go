package rakuten

import (
	"context"
	"fmt"
)

type BooksGenreSearchParams struct {
	BooksGenreID string `url:"booksGenreId,omitempty"`
	GenrePath    int    `url:"genrePath,omitempty"`
}

type BooksGenreSearchResponse struct {
	Current struct {
		BooksGenreID   string `json:"booksGenreId"`
		BooksGenreName string `json:"booksGenreName"`
		GenreLevel     int    `json:"genreLevel"`
	} `json:"current"`
	Children []struct {
		Child struct {
			BooksGenreID   string `json:"booksGenreId"`
			BooksGenreName string `json:"booksGenreName"`
			GenreLevel     int    `json:"genreLevel"`
		} `json:"child"`
	} `json:"children"`
	Parents []interface{} `json:"parents"`
}

func (s *BooksService) GenreSearch(ctx context.Context, opt *BooksGenreSearchParams) (*BooksGenreSearchResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("BooksGenre/Search/20121128?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	respBody := &BooksGenreSearchResponse{}
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}
	return respBody, resp, nil
}
