package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Rules"
)

func testRule(t *testing.T, rule Rules.ValidationRule, fieldName, value string, expectPass bool) {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		passed, msg, _, _ := rule(c, fieldName)
		if passed != expectPass {
			t.Errorf("Value '%s': expected pass=%v, got pass=%v, msg=%s", value, expectPass, passed, msg)
		}
		return nil
	})

	body := bytes.NewBufferString(fieldName + "=" + value)
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	app.Test(req)
}

// تست شماره تلفن
func TestPhoneNumber(t *testing.T) {
	rule := Rules.PhoneNumber()

	testRule(t, rule, "phone", "09123456789", true)  // ✅ معتبر
	testRule(t, rule, "phone", "09351234567", true)  // ✅ معتبر
	testRule(t, rule, "phone", "0912345678", false)  // ❌ کوتاه
	testRule(t, rule, "phone", "08123456789", false) // ❌ پیش‌شماره غلط
}

// تست طول رشته
func TestLengthBetween(t *testing.T) {
	rule := Rules.LengthBetween(3, 10)

	testRule(t, rule, "name", "Ali", true)           // ✅ 3 کاراکتر
	testRule(t, rule, "name", "Mohammad", true)      // ✅ 8 کاراکتر
	testRule(t, rule, "name", "Ab", false)           // ❌ کوتاه
	testRule(t, rule, "name", "VeryLongName", false) // ❌ بلند
}

// تست عدد
func TestNumeric(t *testing.T) {
	rule := Rules.Numeric()

	testRule(t, rule, "price", "100", true)  // ✅
	testRule(t, rule, "price", "0", true)    // ✅
	testRule(t, rule, "price", "abc", false) // ❌
	testRule(t, rule, "price", "-5", false)  // ❌
}

// تست Boolean
func TestBooleanStrict(t *testing.T) {
	rule := Rules.BooleanStrict()

	testRule(t, rule, "active", "true", true)  // ✅
	testRule(t, rule, "active", "false", true) // ✅
	testRule(t, rule, "active", "1", false)    // ❌
	testRule(t, rule, "active", "yes", false)  // ❌
}

// تست Required
func TestRequired(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		value    string
		expected bool
	}{
		{"hello", true},
		{"", false},
	}

	for _, tt := range tests {
		app.Post("/", func(c *fiber.Ctx) error {
			passed, _, _, _ := Rules.Required(c, "field")
			if passed != tt.expected {
				t.Errorf("Required('%s'): expected %v", tt.value, tt.expected)
			}
			return nil
		})

		body := bytes.NewBufferString("field=" + tt.value)
		req := httptest.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		app.Test(req)
	}
}
