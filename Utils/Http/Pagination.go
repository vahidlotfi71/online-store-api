package Http

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginationMetadata struct {
	TotalPages   int  `json:"totalPages"`
	PreviousPage *int `json:"previousPage"`
	NextPage     *int `json:"nextPage"`
	Offset       int  `json:"offset"`
	LimitPerPage int  `json:"limitPerPage"`
	CurrentPage  int  `json:"currentPage"`
}

// Paginate تراکنش ورودی را صفحه‌بندی می‌کند و متادیتا + کوئری آماده را برمی‌گرداند
func Paginate(tx *gorm.DB, c *fiber.Ctx) (*gorm.DB, PaginationMetadata) {
	page, perPage := getPageAndPerPage(c)

	var total int64
	tx.Count(&total)

	meta := calcMetadata(page, perPage, total)
	return tx.Limit(meta.LimitPerPage).Offset(meta.Offset), meta
}

/* ---------- خواندن پارامترها + سقف امن ---------- */
const maxPerPage = 100
const defaultLimit = 15

func getPageAndPerPage(c *fiber.Ctx) (page, perPage int) {
	if p, err := strconv.Atoi(c.Query("page", "1")); err == nil && p > 0 {
		page = p
	} else {
		page = 1
	}

	if pp, err := strconv.Atoi(c.Query("per_page", strconv.Itoa(defaultLimit))); err == nil {
		if pp > maxPerPage {
			pp = maxPerPage
		}
		if pp < 1 {
			pp = defaultLimit
		}
		perPage = pp
	} else {
		perPage = defaultLimit
	}
	return
}

/* ---------- محاسبه متادیتا ---------- */
func calcMetadata(page, perPage int, total int64) PaginationMetadata {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	current := page
	if current > totalPages {
		current = totalPages
	}
	var prev *int
	if current-1 > 0 {
		temp := current - 1
		prev = &temp
	}
	var next *int
	if current+1 <= totalPages {
		temp := current + 1
		next = &temp
	}
	offset := (current - 1) * perPage
	if offset < 0 {
		offset = 0
	}
	return PaginationMetadata{
		TotalPages:   totalPages,
		PreviousPage: prev,
		NextPage:     next,
		Offset:       offset,
		LimitPerPage: perPage,
		CurrentPage:  current,
	}
}
