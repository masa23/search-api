package rakuten

import (
	"context"
	"fmt"
)

type RecipeCategoryParams struct {
	CategoryType string `url:"categoryType,omitempty"`
}

type RecipeCategoryResponse struct {
	Result struct {
		Small []struct {
			CategoryName     string `json:"categoryName"`
			ParentCategoryID string `json:"parentCategoryId"`
			CategoryID       int    `json:"categoryId"`
			CategoryURL      string `json:"categoryUrl"`
		} `json:"small"`
		Medium []struct {
			CategoryName     string `json:"categoryName"`
			ParentCategoryID string `json:"parentCategoryId"`
			CategoryID       int    `json:"categoryId"`
			CategoryURL      string `json:"categoryUrl"`
		} `json:"medium"`
		Large []struct {
			CategoryName string `json:"categoryName"`
			CategoryID   string `json:"categoryId"`
			CategoryURL  string `json:"categoryUrl"`
		} `json:"large"`
	} `json:"result"`
}

func (s RecipeService) CategoryList(ctx context.Context, opt *RecipeCategoryParams) (*RecipeCategoryResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Recipe/CategoryList/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	recipeCategoryResp := &RecipeCategoryResponse{}
	resp, err := s.client.Do(ctx, req, recipeCategoryResp)
	if err != nil {
		return nil, resp, err
	}

	return recipeCategoryResp, resp, nil
}
