package address

import (
	"context"
	"github.com/ekomobile/dadata/v2"
)

type SearchRequest struct {
	Query string `json:"query"`
}

func (s *SearchRequest) getSearchData() (SearchResponse, error) {
	api := dadata.NewCleanApi()
	return api.Address(context.Background(), s.Query)
}
