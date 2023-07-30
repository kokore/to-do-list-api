package queryParams

type OrderQueryParam struct {
	OrderType string `form:"orderType"`
	OrderBy   string `form:"orderBy"`
}

var defaultOrderType string = "DESC"

func (orderQuertParam OrderQueryParam) GetOrderTypeByDefault() string {
	if orderQuertParam.OrderBy != "" {
		if orderQuertParam.OrderType == "" {
			return defaultOrderType
		} else {
			return orderQuertParam.OrderType
		}
	}
	return ""
}
