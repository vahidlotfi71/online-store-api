package rules

import (
	"fmt"
)

func MinLength(n int) Rule {
	return func(v, f string) (bool, string, error) {
		if len(v) < n {
			return false, fmt.Sprintf("%s باید حداقل %d کاراکتر باشد", f, n), nil
		}
		return true, "", nil
	}
}

func MaxLength(n int) Rule {
	return func(v, f string) (bool, string, error) {
		if len(v) > n {
			return false, fmt.Sprintf("%s باید حداکثر %d کاراکتر باشد", f, n), nil
		}
		return true, "", nil
	}
}

func LengthBetween(min, max int) Rule {
	return func(value string, fieldName string) (bool, string, error) {
		length := len(value)
		if length < min || length > max {
			return false, fmt.Sprintf("%s باید بین %d و %d کاراکتر باشد", fieldName, min, max), nil
		}
		return true, "", nil
	}
}
