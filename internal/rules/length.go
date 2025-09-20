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
