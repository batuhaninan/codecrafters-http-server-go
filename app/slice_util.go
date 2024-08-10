package main

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var newSlice []T
	for _, s := range slice {
		if f(s) {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func DeleteEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
