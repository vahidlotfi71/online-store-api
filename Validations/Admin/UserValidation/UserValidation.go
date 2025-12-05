package UserValidation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func Store() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.PhoneNumber()},
		},
		{
			FieldName: "address",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(10, 255)},
		},
		{
			FieldName: "national_id",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.NationalID()},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(8, 16)},
		},
	})
}

func Update() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.PhoneNumber()},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Optional(), Rules.LengthBetween(8, 16)},
		},
	})
}
