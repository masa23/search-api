package rakuten

import (
	"context"
	"fmt"
)

// https://app.rakuten.co.jp/services/api/IchibaGenre/Search/20140222?affiliateId=1a11cd3a.d53d02bd.1a11cd3b.54ce7249&applicationId=1058291654304154113&genreId=0&genrePath=1

type IchibaGenreSearchParams struct {
	GenreID   int `url:"genreId,omitempty"`
	GenrePath int `url:"genrePath,omitempty"`
}

type IchibaGenreResponse struct {
	Parents []interface{} `json:"parents"`
	Current struct {
		GenreLevel int    `json:"genreLevel"`
		GenreName  string `json:"genreName"`
		GenreID    int    `json:"genreId"`
	} `json:"current"`
	Children []struct {
		Child struct {
			GenreLevel int    `json:"genreLevel"`
			GenreName  string `json:"genreName"`
			GenreID    int    `json:"genreId"`
		} `json:"child"`
	} `json:"children"`
	TagGroups []interface{} `json:"tagGroups"`
	Brothers  []interface{} `json:"brothers"`
}

func (s IchibaService) GenreSearch(ctx context.Context, opt *IchibaGenreSearchParams) (*IchibaGenreResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("IchibaGenre/Search/20140222?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	ichibaGenreResp := &IchibaGenreResponse{}
	resp, err := s.client.Do(ctx, req, ichibaGenreResp)
	if err != nil {
		return nil, resp, err
	}

	return ichibaGenreResp, resp, nil
}
