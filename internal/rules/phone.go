package rules

import (
	"fmt"
	"regexp"
)

func Phone() Rule {
	re := regexp.MustCompile(`^09[0-9]{9}$`)
	return func(v, f string) (bool, string, error) {
		if !re.MatchString(v) {
			return false, fmt.Sprintf("%s برای شماره موبایل معتبر باشد", f), nil
		}
		return true, "", nil
	}
}
