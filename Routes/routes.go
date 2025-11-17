package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Controllers/AdminController"
	"github.com/vahidlotfi71/online-store-api/Controllers/AuthController"
	"github.com/vahidlotfi71/online-store-api/Controllers/OrderController"
	"github.com/vahidlotfi71/online-store-api/Controllers/ProductController"
	"github.com/vahidlotfi71/online-store-api/Controllers/UserController"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	adminProductVal "github.com/vahidlotfi71/online-store-api/Validations/Admin/ProductValidation"
	adminUserVal "github.com/vahidlotfi71/online-store-api/Validations/Admin/UserValidation"
	authVal "github.com/vahidlotfi71/online-store-api/Validations/Auth"
	userVal "github.com/vahidlotfi71/online-store-api/Validations/User"
)

func SetupRoutes(app *fiber.App) {

	// ---------- PUBLIC AUTH ----------
	app.Post("/admin/login", authVal.AdminLogin(), AuthController.AdminLogin)
	app.Post("/login", authVal.Login(), AuthController.Login)
	app.Post("/register", authVal.Register(), AuthController.Register)
	app.Post("/logout", AuthController.AdminLogout) // می‌توانید نسخه‌ی جدا برای user هم بسازید

	// ---------- ADMIN GROUP ----------
	admin := app.Group("/admin", Middlewares.AdminAuthMiddleware)

	// dashboard
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Admin Dashboard"})
	})

	// ---------- ADMIN.USER ----------
	adminUser := admin.Group("/user")
	adminUser.Get("/", AdminController.Index).Name("admin.user.index")
	adminUser.Get("/show/:id", AdminController.Show).Name("admin.user.show")
	adminUser.Post("/store", adminUserVal.Store(), AdminController.Store).Name("admin.user.store")
	adminUser.Post("/update/:id", adminUserVal.Update(), AdminController.Update).Name("admin.user.update")
	adminUser.Post("/delete/:id", AdminController.Delete).Name("admin.user.delete")
	adminUser.Get("/restore/:id", AdminController.Restore).Name("admin.user.restore")
	adminUser.Get("/trash", AdminController.Trash).Name("admin.user.trash")
	adminUser.Post("/clear-trash", AdminController.ClearTrash).Name("admin.user.clear-trash")

	// ---------- ADMIN.PRODUCT ----------
	adminProduct := admin.Group("/product")
	adminProduct.Get("/", ProductController.Index).Name("admin.product.index")
	adminProduct.Get("/show/:id", ProductController.Show).Name("admin.product.show")
	adminProduct.Post("/store", adminProductVal.CreateProduct(), ProductController.Store).Name("admin.product.store")
	adminProduct.Post("/update/:id", adminProductVal.UpdateProduct(), ProductController.Update).Name("admin.product.update")
	adminProduct.Post("/delete/:id", ProductController.Delete).Name("admin.product.delete")
	adminProduct.Get("/restore/:id", ProductController.Restore).Name("admin.product.restore")
	adminProduct.Get("/trash", ProductController.Trash).Name("admin.product.trash")
	adminProduct.Post("/clear-trash", ProductController.ClearTrash).Name("admin.product.clear-trash")

	// ---------- ADMIN.ORDER ----------
	adminOrder := admin.Group("/order")
	adminOrder.Get("/", OrderController.Index).Name("admin.order.index")
	adminOrder.Get("/show/:id", OrderController.Show).Name("admin.order.show")
	adminOrder.Post("/update/:id", OrderController.Update).Name("admin.order.update")
	adminOrder.Get("/trash", OrderController.Trash).Name("admin.order.trash")

	// ---------- USER (AUTHENTICATED) ----------
	user := app.Group("/user", Middlewares.AuthMiddleware)
	user.Get("/profile", UserController.Show)                                    // نمایش پروفایل خود کاربر
	user.Post("/profile/update", userVal.Update(), UserController.Update)        // ویرایش پروفایل
	user.Get("/orders", OrderController.Index)                                   // لیست سفارشات خود کاربر
	user.Post("/orders", userVal.CreateOrderValidation(), OrderController.Store) // ثبت سفارش جدید

	// ---------- PUBLIC PRODUCT ----------
	app.Get("/products", ProductController.Index)    // لیست محصولات فعال
	app.Get("/products/:id", ProductController.Show) // مشخصات یک محصول

	// ---------- 404 HANDLER ----------
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Route Not Found"})
	})
}
