package rules

type Rule func(value string, fieldName string) (passed bool, message string, err error)

type FieldRules struct {
	Field string `json:"field"`
	Rules []Rule `json:"rules"`
}

type ValidationError struct {
	Field   string `json:"field_name"`
	Message string `json:"message"`
}
