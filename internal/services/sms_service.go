package services

import (
	"fmt"

	"github.com/vahidlotfi71/online-store-api.git/config"
)

type SMSService struct {
	CFG *config.Config
}

// یک ساختار (struct) به اسم SMSService میسازیم  که تنظیمات (Config) رو نگه می‌داره.
// این یعنی وقتی می‌خواهیم پیامک بفرستیم، باید بدانیم تنظیمات SMS Gateway چی هست.
func NewSMSService(cfg *config.Config) *SMSService { return &SMSService{CFG: cfg} }

// این یک constructor هست.
//
//	اجازه میده با دادن config یک نمونه‌ی جدید از SMSService بسازیم
func (s *SMSService) SendSMS(to, msg string) error {
	// در حالت توسعه فقط لاگ می‌کنیم
	fmt.Printf("SMS to %s: %s\n", to, msg)
	// در حالت واقعی اینجا به کاوه‌نگار متصل می‌شویم
	return nil
}
