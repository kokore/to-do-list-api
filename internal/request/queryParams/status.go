package queryParams

type StatusQueryParam struct {
	StatusQuery *string `form:"statusType"`
}

var defaultStatusType string = "DESC"

func (statusQueryParam StatusQueryParam) GetStatusTypeByDefault() *string {
	if statusQueryParam.StatusQuery == nil {
		return &defaultStatusType
	} else {
		return statusQueryParam.StatusQuery
	}
}
