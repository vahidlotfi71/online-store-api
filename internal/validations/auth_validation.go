package validations

import (
	"github.com/vahidlotfi71/online-store-api.git/internal/rules"
)

func RegisterValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{
			Field: "first_name",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MinLength(2),
				rules.MaxLength(50),
			},
		},
		{
			Field: "last_name",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MinLength(2),
				rules.MaxLength(50),
			},
		},
		{
			Field: "phone",
			Rules: []rules.Rule{
				rules.Required(),
				rules.Phone(),
			},
		},
		{
			Field: "address",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MaxLength(500),
			},
		},
		{Field: "national_ID",
			Rules: []rules.Rule{
				rules.Required(),
				rules.NationalID(),
			},
		},
		{Field: "password",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MinLength(8),
			},
		},
	}
}

func LoginValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{
			Field: "phone",
			Rules: []rules.Rule{
				rules.Required(),
				rules.Phone(),
			},
		},
		{Field: "password",
			Rules: []rules.Rule{
				rules.Required(),
			},
		},
	}
}

func VerifyPhoneValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{Field: "phone", Rules: []rules.Rule{rules.Required(), rules.Phone()}},
		{Field: "code", Rules: []rules.Rule{rules.Required(), rules.LengthBetween(6, 6)}},
	}
}
