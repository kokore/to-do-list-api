package queryParams

type TitleQueryParam struct {
	TitleQuery *string `form:"titleType"`
}

var defaultTitleType string = "DESC"

func (titleQueryParam TitleQueryParam) GetTitleTypeByDefault() *string {
	if titleQueryParam.TitleQuery == nil {
		return &defaultTitleType
	} else {
		return titleQueryParam.TitleQuery
	}
}
