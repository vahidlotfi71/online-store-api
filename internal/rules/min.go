package rules

import (
	"fmt"
	"strconv"
)

func Min(min float64) Rule {
	return func(v, f string) (bool, string, error) {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return false, fmt.Sprintf("%s باید عدد باشد", f), nil
		}
		if n < min {
			return false, fmt.Sprintf("%s باید حداقل %v باشد", f, min), nil
		}
		return true, "", nil

	}
}
