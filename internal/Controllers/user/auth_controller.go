package user

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/online-store-api.git/internal/models"
	"github.com/vahidlotfi71/online-store-api.git/internal/services"
	"github.com/vahidlotfi71/online-store-api.git/internal/utils"
)

type AuthController struct {
	AuthService *services.AuthService
	SMSService  *services.SMSService
}

func NewAuthController(as *services.AuthService, sms *services.SMSService) *AuthController {
	return &AuthController{AuthService: as, SMSService: sms}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	// یک ساختار موقت می سازیم تا اطلاعات ورودی ثبت‌نام (نام، شماره، پسورد و ...) گرفته بشه.
	var input struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Phone      string `json:"phone"`
		Address    string `json:"address"`
		NationalID string `json:"national_id"`
		Password   string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The data format is incorrect.")
	}

	// بررسی می‌کنیم که این شماره موبایل قبلاً ثبت نشده باشه.
	if _, err := ac.AuthService.GetUserByPhone(input.Phone); err == nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, "This phone number is already registered.")
	}

	// پسورد هش میشه
	hashedPass, _ := utils.HashPassword(input.Password)
	// یک کد ۶ رقمی تصادفی تولید میشه (کد تایید).
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// یک کاربر جدید ساخته میشه
	user := &models.User{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Phone:      input.Phone,
		Address:    input.Address,
		NationalID: input.NationalID,
		Password:   hashedPass,
		Role:       "user",
		IsVerified: false,
		VerifyCode: code,
	}

	// کاربر در دیتابیس ذخیره میشه.
	if err := ac.AuthService.CreateUser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error registering data")
	}

	// پیامک کد تایید برای کاربر فرستاده میشه.
	msg := fmt.Sprintf("Your verification code: %s", code)
	ac.SMSService.SendSMS(input.Phone, msg)

	// موفقیت ثبت‌نام و user_id برگردونده میشه.
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Registration completed successfully",
		"data": fiber.Map{
			"user_id": user.ID,
		},
	})
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	// ورودی مورد نیاز: شماره موبایل + پسورد.
	var input struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The data format is incorrect.")
	}

	// کاربر از دیتابیس پیدا میشه.
	// پسورد بررسی میشه (مقایسه با هش).
	user, err := ac.AuthService.GetUserByPhone(input.Phone)
	if err != nil || !utils.CheckPassword(input.Password, user.Password) {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "The phone number or password is incorrect.")
	}

	// اگر شماره تایید نشده باشه → اجازه ورود نداره.
	if !user.IsVerified {
		return utils.ErrorResponse(c, fiber.StatusForbidden, "Your phone number is not verified.")
	}
	// اگر همه‌چیز درست باشه → توکن JWT ساخته میشه.
	token, err := utils.GenerateJWT(user.ID, user.Phone, user.Role, ac.AuthService.CFG)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(fiber.Map{"success": true, "message": "Login successful", "data": fiber.Map{"token": token, "user": user}})
}

// شماره و کد را می‌گیریم.
// توی دیتابیس دنبال کاربر می‌گردیم.
// کد رو چک می‌کنیم.
// کاربر رو فعال می‌کنیم.
// براش یک توکن JWT می‌سازیم.
// جواب JSON برمی‌گردونیم.
func (ac *AuthController) VerifyPhone(c *fiber.Ctx) error {
	var input struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The data format is incorrect.")
	}
	user, err := ac.AuthService.GetUserByPhone(input.Phone)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}
	if user.VerifyCode != input.Code {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "The verification code is incorrect.")
	}
	user.IsVerified = true
	user.VerifyCode = ""
	if err := ac.AuthService.UpdateUser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Account activation error")
	}
	token, _ := utils.GenerateJWT(user.ID, user.Phone, user.Role, ac.AuthService.CFG)
	return c.JSON(fiber.Map{"success": true, "message": "Account activated successfully", "data": fiber.Map{"token": token, "user": user}})
}

func (ac *AuthController) GetProfile(c *fiber.Ctx) error {
	// userID از توکن JWT (که قبلاً در میدل‌ور ذخیره شده) گرفته میشه.
	userID := c.Locals("userID").(uint)

	user, err := ac.AuthService.GetUserByID(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}

	return utils.SuccessResponse(c, user)
}
