package Auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func Register() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(3, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(3, 255)},
		},
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.PhoneNumber()},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(8, 16)},
		},
		{
			FieldName: "address",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(10, 300)},
		},
		{
			FieldName: "national_ID",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.NationalID()},
		},
		// فیلد role حذف شد چون در منطق بیزینس مقداردهی می‌شود
	})
}
