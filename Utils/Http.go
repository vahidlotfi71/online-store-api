package Utils

import "github.com/gofiber/fiber/v2"

// این ساختار برای نمایش خطاهای جزئی مانند خطاهای اعتبارسنجی استفاده می‌شود.
type ApiError struct {
	Field   string `json:"field,omitempty"` // فقط برای خطاهای ولیدیشن
	Message string `json:"message"`
}

// این ساختار، قالب یکپارچه پاسخ‌های API (پیام، داده و خطاها) را مشخص می‌کند.
type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []ApiError  `json:"errors,omitempty"`
}

// این تابع کمکی، یک پاسخ JSON سازگار با ساختار ApiResponse تولید می‌کند.
// status: کد وضعیت HTTP (مانند 200, 201, 400, 500)
// message: پیامی برای کاربر در مورد عملیات انجام شده.
// data: محتوای اصلی پاسخ (مثلاً اطلاعات کاربر یا لیست محصولات).
// errors: آرایه‌ای از خطاهای دقیق (مانانند خطاهای ولیدیشن).
func Response(c *fiber.Ctx, status int, message string, data interface{}, errors []ApiError) error {
	return c.Status(status).JSON(ApiResponse{
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}

// SimpleSuccess is a shortcut for a successful response (200 OK).
func SimpleSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return Response(c, 200, message, data, nil)
}

// SimpleError is a shortcut for a generic error response (usually 400 or 500).
func SimpleError(c *fiber.Ctx, status int, message string) error {
	return Response(c, status, message, nil, nil)
}
