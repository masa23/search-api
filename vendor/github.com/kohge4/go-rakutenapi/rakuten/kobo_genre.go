package rakuten

import (
	"context"
	"fmt"
)

type KoboGenreParams struct {
	KoboGenreID int `url:"koboGenreId,omitempty"`
	GenrePath   int `url:"genrePath,omitempty"`
}

type KoboGenreResponse struct {
	Current struct {
		KoboGenreID   string `json:"koboGenreId"`
		KoboGenreName string `json:"koboGenreName"`
		GenreLevel    int    `json:"genreLevel"`
	} `json:"current"`
	Children []struct {
		Child struct {
			KoboGenreID   string `json:"koboGenreId"`
			KoboGenreName string `json:"koboGenreName"`
			GenreLevel    int    `json:"genreLevel"`
		} `json:"child"`
	} `json:"children"`
	Parents []interface{} `json:"parents"`
}

func (s KoboService) GenreSearch(ctx context.Context, opt *KoboGenreParams) (*KoboGenreResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Kobo/GenreSearch/20131010?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	koboGenreResp := &KoboGenreResponse{}
	resp, err := s.client.Do(ctx, req, koboGenreResp)
	if err != nil {
		return nil, resp, err
	}

	return koboGenreResp, resp, nil
}
