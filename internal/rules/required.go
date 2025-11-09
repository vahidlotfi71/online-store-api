package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Required(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return false, "", nil, err
	}

	value := form.Value[field_name]
	file := form.File[field_name]

	message = fmt.Sprintf("The %s field is required", field_name)

	if (value == nil || len(value) == 0 || value[0] == "") && (file == nil || len(file) == 0) {
		return false, message, nil, nil
	}

	return true, "", nil, nil
}
