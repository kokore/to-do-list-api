package queryParams

type SearchQueryParam struct {
	SearchText *string `form:"search"`
}

func (searchQueryParam SearchQueryParam) HasSearchText() bool {
	return !(searchQueryParam.SearchText == nil || *searchQueryParam.SearchText == "")
}

func (searchQueryParam SearchQueryParam) GetSearchText() string {
	return *searchQueryParam.SearchText
}
