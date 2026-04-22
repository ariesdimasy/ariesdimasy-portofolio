package utils

import "strconv"

// StrToInt converts a string to int. Returns defaultValue if the string is empty or invalid.
func StrToInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return v
}

func StringToUint(s string) (uint, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}
