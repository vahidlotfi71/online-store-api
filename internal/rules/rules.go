package rules

type Rule func(value string, fieldName string) (passed bool, message string, err error)

type FieldRules struct {
	FieldName string `json:"field"`
	Rules     []Rule `json:"rules"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
