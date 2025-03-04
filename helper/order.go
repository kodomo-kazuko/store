package helper

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func OrderBy(orderCol field.OrderExpr, c *fiber.Ctx, order field.OrderExpr) func(db gen.Dao) gen.Dao {
	return func(db gen.Dao) gen.Dao {
		if orderCol == nil {
			if order != nil {
				return db.Order(order.Desc())
			}
			return db
		}

		direction := c.Query("sort_type")
		switch direction {
		case "desc":
			return db.Order(orderCol.Desc())
		case "asc":
			return db.Order(orderCol.Asc())
		default:
			return db.Order(orderCol.Asc())
		}
	}
}
