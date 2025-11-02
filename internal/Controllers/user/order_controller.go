package user

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"github.com/vahidlotfi71/online-store-api.git/internal/services"
	"github.com/vahidlotfi71/online-store-api.git/internal/utils"
)

type OrderController struct {
	OrderService   *services.OrderService   //OrderService → برای کارهای مربوط به سفارش (ثبت، دریافت، حذف و...)
	ProductService *services.ProductService //ProductService → برای دسترسی به محصولات (مثلاً بررسی موجودی)
	SMSService     *services.SMSService     //SMSService → برای ارسال پیامک تأیید خرید
	AuthService    services.AuthService
}

func NewOrderController(os *services.OrderService, ps *services.ProductService, sms *services.SMSService) *OrderController {
	//This function is the standard way to create a new instance of the controller.
	return &OrderController{
		OrderService:   os,
		ProductService: ps,
		SMSService:     sms,
	}
}

// Get products list
func (oc *OrderController) GetProducts(c *fiber.Ctx) error {
	page, limit := 1, 10
	//c.Query("page") و c.Query("limit") یعنی اگر کاربر از طریق URL صفحه خاصی خواست، آن را بخوان
	if p := c.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}

	if l := c.Query("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
	}

	filter := map[string]interface{}{}
	if name := c.Query("name"); name != "" {
		filter["name"] = name
	}
	if brand := c.Query("brand"); brand != "" {
		filter["brand"] = brand
	}

	list, total, err := oc.ProductService.GetProducts(filter, page, limit)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "خطا در دریافت اطلاعات")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "لیست محصولات",
		"data": fiber.Map{
			"products": list,
			"pagination": fiber.Map{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	},
	)
}

func (oc *OrderController) CreateOrder(c *fiber.Ctx) error {
	//get UserID
	userID := c.Locals("userID").(uint)

	//Definition of input structure
	var input struct {
		Items []models.OrderItem `json:"items"` // [{product_id:1,quantity:2}, ...]
	}

	// c.BodyParser(&input) reads the JSON body data and maps it into the input struct.
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "فرمت داده‌ها نادرست است")
	}

	order, err := oc.OrderService.CreateOrder(userID, input.Items)
	if err != nil {
		// We used err.Error() because we don't know exactly what error might occur.

		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, _ := oc.AuthService.GetUserByID(userID)

	msg := fmt.Sprintf("Hello %s %s,\nYour purchase was successful.\nProducts: %d item(s)\nTotal price: %d Tomans\nThank you, Our Store",
		user.FirstName,
		user.LastName,
		len(order.Items),
		int(order.TotalPrice))
	oc.SMSService.SendSMS(user.Phone, msg)

	return c.JSON(fiber.Map{"success": true, "message": "Purchase completed successfully", "data": order})
}

func (oc *OrderController) GetUserOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	list, err := oc.OrderService.GetUserOrders(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error retrieving data")
	}
	return utils.SuccessResponse(c, list)
}
