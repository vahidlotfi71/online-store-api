package rules

import (
	"fmt"
)

func MinLength(n int) Rule {
	return func(v, f string) (bool, string, error) {
		if len(v) < n {
			return false, fmt.Sprintf("%s must be at least %d characters", f, n), nil
		}
		return true, "", nil
	}
}

func MaxLength(n int) Rule {
	return func(v, f string) (bool, string, error) {
		if len(v) > n {
			return false, fmt.Sprintf("%s must be at most %d characters", f, n), nil
		}
		return true, "", nil
	}
}

func LengthBetween(min, max int) Rule {
	return func(value string, fieldName string) (bool, string, error) {
		length := len(value)
		if length < min || length > max {
			return false, fmt.Sprintf("%s must be between %d and %d characters", fieldName, min, max), nil
		}
		return true, "", nil
	}
}
