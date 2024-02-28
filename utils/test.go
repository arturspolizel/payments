package utils

import "reflect"

// CheckTestError uses reflection to compare error types in places (table tests, for example)
// where errors.As cannot be used easily. It does not support error wrapping, so use with caution.
func CheckTestError(got, want error) bool {
	return reflect.TypeOf(got) == reflect.TypeOf(want)
}
