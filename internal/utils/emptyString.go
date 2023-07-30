package utils

func IsEmptyString(value *string) bool {
	if value == nil {
		return true
	}
	if *value == "" {
		return true
	}

	return false
}
