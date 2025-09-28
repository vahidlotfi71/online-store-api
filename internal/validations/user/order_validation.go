package user

import "github.com/vahidlotfi71/online-store-api.git/internal/rules"

func CreateOrderValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{
			Field: "items",
			Rules: []rules.Rule{
				rules.Required(),
			},
		},
	}
}

func UpdateOrderValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{
			Field: "items",
			Rules: []rules.Rule{
				rules.Required(),
			},
		},
	}
}
