package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) func(db gen.Dao) gen.Dao {

	return func(db gen.Dao) gen.Dao {
		q := c.Query("page_number")
		page, _ := strconv.Atoi(q)

		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))

		switch {
		case pageSize >= 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateQ(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := c.Query("page_number")
		page, _ := strconv.Atoi(q)
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize >= 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateR(c *fiber.Ctx) (int, int) {
	var pageNumber int
	var pageSize int
	pageNumber, _ = strconv.Atoi(c.Query("page_number"))
	pageSize, _ = strconv.Atoi(c.Query("page_size"))
	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	offset := (pageNumber - 1) * pageSize

	return offset, pageSize
}
