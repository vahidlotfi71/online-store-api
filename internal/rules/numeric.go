package rules

import (
	"fmt"
	"regexp"
)

func Numeric() Rule {
	return func(v, f string) (bool, string, error) {
		if matched, _ := regexp.MatchString(`^\d+$`, v); !matched {
			return false, fmt.Sprintf("%s باید عدد باشد", f), nil
		}
		return true, "", nil
	}

}
