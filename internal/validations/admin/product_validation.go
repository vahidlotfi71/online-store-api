package admin

import "github.com/vahidlotfi71/online-store-api.git/internal/rules"

var CreateProductValidation = createProductValidation
var UpdateProductValidation = createProductValidation // اگر متفاوت بود جدا می کنیم

func createProductValidation() []rules.FieldRules {
	return []rules.FieldRules{
		{
			Field: "name",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MinLength(2),
				rules.MaxLength(100),
			},
		},

		{
			Field: "brand",
			Rules: []rules.Rule{
				rules.Required(),
				rules.MinLength(2),
				rules.MaxLength(50),
			},
		},
		{
			Field: "price",
			Rules: []rules.Rule{
				rules.Required(),
				rules.Numeric(),
				rules.Min(0),
			},
		},
		{
			Field: "stock",
			Rules: []rules.Rule{
				rules.Required(),
				rules.Numeric(),
				rules.Min(0),
			},
		},
		{
			Field: "description",
			Rules: []rules.Rule{
				rules.MaxLength(1000),
			},
		},
	}
}
