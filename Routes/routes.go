package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api/Controllers/AuthController"
	"github.com/vahidlotfi71/online-store-api/Controllers/OrderController"
	"github.com/vahidlotfi71/online-store-api/Controllers/ProductController"
	"github.com/vahidlotfi71/online-store-api/Controllers/UserController"
	"github.com/vahidlotfi71/online-store-api/Middlewares"
	"github.com/vahidlotfi71/online-store-api/Models"
	"github.com/vahidlotfi71/online-store-api/Resources/UserResource"
	"github.com/vahidlotfi71/online-store-api/Validations/Admin/ProductValidation"
	"github.com/vahidlotfi71/online-store-api/Validations/Admin/UserValidation"
	"github.com/vahidlotfi71/online-store-api/Validations/Auth"
)

func SetupRoutes(app *fiber.App) {

	// ---------- PUBLIC AUTH ----------
	app.Post("/admin/login", Auth.AdminLogin(), AuthController.AdminLogin)
	app.Post("/login", Auth.Login(), AuthController.Login)
	app.Post("/register", Auth.Register(), AuthController.Register)
	app.Post("/logout", AuthController.AdminLogout)

	// ---------- ADMIN GROUP ----------
	admin := app.Group("/admin", Middlewares.AdminAuthMiddleware)

	// dashboard
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Admin Dashboard"})
	})

	// ---------- ADMIN.USER ----------
	adminUser := admin.Group("/user")
	adminUser.Get("/", UserController.Index).Name("admin.user.index")
	adminUser.Get("/show/:id", UserController.Show).Name("admin.user.show")
	adminUser.Post("/store", UserValidation.Store(), UserController.Store).Name("admin.user.store")
	adminUser.Post("/update/:id", UserValidation.Update(), UserController.Update).Name("admin.user.update")
	adminUser.Post("/delete/:id", UserController.Delete).Name("admin.user.delete")
	adminUser.Get("/restore/:id", UserController.Restore).Name("admin.user.restore")
	adminUser.Get("/trash", UserController.Trash).Name("admin.user.trash")
	adminUser.Post("/clear-trash", UserController.ClearTrash).Name("admin.user.clear-trash")

	// ---------- ADMIN.PRODUCT ----------
	adminProduct := admin.Group("/product")
	adminProduct.Get("/", ProductController.Index).Name("admin.product.index")
	adminProduct.Get("/show/:id", ProductController.Show).Name("admin.product.show")
	adminProduct.Post("/store", ProductValidation.CreateProduct(), ProductController.Store).Name("admin.product.store")
	adminProduct.Post("/update/:id", ProductValidation.UpdateProduct(), ProductController.Update).Name("admin.product.update")
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
	user := app.Group("/user", Middlewares.UserAuthMiddleware)
	user.Get("/profile", func(c *fiber.Ctx) error {
		user := c.Locals("user").(Models.User)
		return c.JSON(fiber.Map{"data": UserResource.Single(user)})
	})
	user.Post("/profile/update", UserValidation.Update(), UserController.Update)
	user.Get("/orders", OrderController.Index)
	// user.Post("/orders", OrderController.Store) // بعداً اضافه شود

	// ---------- PUBLIC PRODUCT ----------
	app.Get("/products", ProductController.Index)
	app.Get("/products/:id", ProductController.Show)

	// ---------- 404 HANDLER ----------
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Route Not Found"})
	})
}
