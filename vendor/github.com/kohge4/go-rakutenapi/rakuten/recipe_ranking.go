package rakuten

import (
	"context"
	"fmt"
)

type RecipeRankingParams struct {
	CategoryID string `url:"categoryId,omitempty"`
}

type RecipeRankingResponse struct {
	Result []struct {
		FoodImageURL      string   `json:"foodImageUrl"`
		RecipeDescription string   `json:"recipeDescription"`
		RecipePublishday  string   `json:"recipePublishday"`
		Shop              int      `json:"shop"`
		Pickup            int      `json:"pickup"`
		RecipeID          int      `json:"recipeId"`
		Nickname          string   `json:"nickname"`
		SmallImageURL     string   `json:"smallImageUrl"`
		RecipeMaterial    []string `json:"recipeMaterial"`
		RecipeIndication  string   `json:"recipeIndication"`
		RecipeCost        string   `json:"recipeCost"`
		Rank              string   `json:"rank"`
		RecipeURL         string   `json:"recipeUrl"`
		MediumImageURL    string   `json:"mediumImageUrl"`
		RecipeTitle       string   `json:"recipeTitle"`
	} `json:"result"`
}

func (s RecipeService) Ranking(ctx context.Context, opt *RecipeRankingParams) (*RecipeRankingResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("Recipe/CategoryRanking/20170426?")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	recipeRankingResp := &RecipeRankingResponse{}
	resp, err := s.client.Do(ctx, req, recipeRankingResp)
	if err != nil {
		return nil, resp, err
	}

	return recipeRankingResp, resp, nil
}
