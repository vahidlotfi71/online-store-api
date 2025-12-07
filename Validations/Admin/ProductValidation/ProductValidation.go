package ProductValidation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func CreateProduct() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "brand",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "price",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		},
		{
			FieldName: "description",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(10, 255)},
		},
		{
			FieldName: "stock",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Numeric()},
		}, {
			FieldName: "is_active",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.BooleanStrict()},
		},
	})
}

// فایل: Validations/Admin/ProductValidation/ProductValidation.go

func UpdateProduct() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "name",
			// اگر 'name' فرستاده نشود، Optional اجرا شده و بقیه قوانین نادیده گرفته می‌شوند
			Rules: []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "brand",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(2, 100)},
		},
		{
			FieldName: "price",
			Rules:     []Rules.ValidationRule{Rules.Numeric()},
		},
		{
			FieldName: "description",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(10, 255)},
		},
		{
			FieldName: "stock",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.Numeric()},
		},
		{
			FieldName: "is_active",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.BooleanStrict()},
		},
		// فیلدهای دیگر...
	})
}
