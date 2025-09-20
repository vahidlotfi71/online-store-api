package rules

import (
	"regexp"
	"strconv"
)

func NationalID() Rule {
	return func(v, f string) (bool, string, error) {
		if matched, _ := regexp.MatchString(`^\d{10}$`, v); !matched {
			return false, "کد ملی باید ۱۰ رقم باشد", nil
		}
		sum := 0
		for i := 0; i < 9; i++ {
			d, _ := strconv.Atoi(string(v[i]))
			sum += d * (10 - i)
		}
		remainder := sum % 11
		check, _ := strconv.Atoi(string(v[9]))
		if remainder < 2 {
			if check != remainder {
				return false, "کد ملی معتبر نیست", nil
			}
		} else {
			if check != (11 - remainder) {
				return false, "کد ملی معتبر نیست", nil
			}
		}
		return true, "", nil
	}
}
