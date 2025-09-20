package rules

import "fmt"

func Required() Rule {
	return func(v, f string) (bool, string, error) {
		if v == "" {
			return false, fmt.Sprintf("فیلد %s اجباری است", f), nil
		}
		return true, "", nil
	}

}
