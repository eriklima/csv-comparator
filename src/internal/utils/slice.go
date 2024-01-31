package utils

func SliceFilter[T any](slice []T, criteria func(T) bool) (result []T) {
	for _, value := range slice {
		if criteria(value) {
			result = append(result, value)
		}
	}
	return
}