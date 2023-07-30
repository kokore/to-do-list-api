package queryParams

type DateQueryParam struct {
	DateQuery *string `form:"dateType"`
}

var defaultDateType string = "DESC"

func (dateQueryParam DateQueryParam) GetDateTypeByDefault() *string {
	if dateQueryParam.DateQuery == nil {
		return &defaultDateType
	} else {
		return dateQueryParam.DateQuery
	}
}
