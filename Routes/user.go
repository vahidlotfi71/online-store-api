package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/config"
	"github.com/vahidlotfi71/online-store-api.git/internal/Controllers/user"
	"github.com/vahidlotfi71/online-store-api.git/internal/middlewares"
	"github.com/vahidlotfi71/online-store-api.git/internal/services"
	"github.com/vahidlotfi71/online-store-api.git/internal/validations"
	userVal "github.com/vahidlotfi71/online-store-api.git/internal/validations/user"

	"gorm.io/gorm"
)

func setupUserRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	as := services.NewAuthService(db, cfg)
	sms := services.NewSMSService(cfg)
	ps := services.NewProductService(db)
	os := services.NewOrderService(db)

	authCtrl := user.NewAuthController(as, sms)
	orderCtrl := user.NewOrderController(os, ps, sms)

	// Definition of public routes (no login required)
	api.Post("/register", //register → User registration
		middlewares.ValidationMiddleware(validations.RegisterValidation()),
		authCtrl.Register)
	api.Post("/login", //login → User login
		middlewares.ValidationMiddleware(validations.LoginValidation()),
		authCtrl.Login,
	)
	api.Post("/verify", //verify → Phone number verification
		middlewares.ValidationMiddleware(validations.VerifyPhoneValidation()),
		authCtrl.VerifyPhone)

	// Definition of routes for logged-in users
	user := api.Group("/user", middlewares.AuthMiddleware(cfg))
	user.Get("/profile", authCtrl.GetProfile)
	user.Get("/products", orderCtrl.GetProducts)
	// ایجاد سفارش جدید (قبل از آن ورودی‌ها اعتبارسنجی می‌شن)
	user.Post("/orders", middlewares.ValidationMiddleware(userVal.CreateOrderValidation()), orderCtrl.CreateOrder)
	user.Get("/orders", orderCtrl.GetUserOrders)
}
