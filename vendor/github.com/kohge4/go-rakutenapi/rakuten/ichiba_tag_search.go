package rakuten

import (
	"context"
	"fmt"
)

type IchibaTagParams struct {
	TagID int `url:"tagId,omitempty"`
}

type IchibaTagResponse struct {
	TagGroups []struct {
		TagGroup struct {
			Tags []struct {
				Tag struct {
					ParentTagID int    `json:"parentTagId"`
					TagName     string `json:"tagName"`
					TagID       int    `json:"tagId"`
				} `json:"tag"`
			} `json:"tags"`
			TagGroupID   int    `json:"tagGroupId"`
			TagGroupName string `json:"tagGroupName"`
		} `json:"tagGroup"`
	} `json:"tagGroups"`
}

func (s *IchibaService) TagSearch(ctx context.Context, opt *IchibaTagParams) (*IchibaTagResponse, *Response, error) {
	urlSuffix := fmt.Sprintf("IchibaTag/Search/20140222")

	req, err := s.client.NewRequest("GET", urlSuffix, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	ichibaTagResp := &IchibaTagResponse{}
	resp, err := s.client.Do(ctx, req, ichibaTagResp)
	if err != nil {
		return nil, resp, err
	}

	return ichibaTagResp, resp, nil
}
