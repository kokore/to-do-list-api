package utils

func SortConditions(value string) int {
	if value == "DESC" {
		return -1
	}
	return 1
}
