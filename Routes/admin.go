package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/config"
	"github.com/vahidlotfi71/online-store-api.git/internal/Controllers/admin"
	"github.com/vahidlotfi71/online-store-api.git/internal/middlewares"
	"github.com/vahidlotfi71/online-store-api.git/internal/services"
	adminVal "github.com/vahidlotfi71/online-store-api.git/internal/validations/admin"
	"gorm.io/gorm"
)

func setupAdminRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	ps := services.NewProductService(db)
	productCtrl := admin.NewProductController(ps)

	admin := api.Group("/admin", middlewares.AuthMiddleware(cfg), middlewares.RequireRole("admin"))

	// create product
	admin.Post("/products",
		middlewares.ValidationMiddleware(adminVal.CreateProductValidation()),
		productCtrl.CreateProduct)

	//get all products
	admin.Get("/products",
		productCtrl.GetProducts)

	//get product by id
	admin.Get("/products/:id",
		productCtrl.GetProductByID)

	//update product
	// We used Post instead of Put because they have the same functionality.
	admin.Post("/products/:id",
		// The submitted data is first validated, then sent to the controller.
		middlewares.ValidationMiddleware(adminVal.UpdateProductValidation()),
		productCtrl.UpdateProduct)

	// delete product
	admin.Delete("/products/:id", productCtrl.DeleteProduct)

}
