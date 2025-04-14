package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gen"
)

type QueryInput func(string) gen.Condition

type QueryType map[string]QueryInput

// Constructor to wrap a function as QueryInput
func QueryFunc(f func(string) gen.Condition) QueryInput {
	return f
}
func Where(conds ...gen.Condition) func(db gen.Dao) gen.Dao {
	return func(db gen.Dao) gen.Dao {
		if conds == nil {
			return db
		}
		return db.Where(conds...)
	}
}

func BuildConds(c *fiber.Ctx, paramToCondition map[string]QueryInput) []gen.Condition {
	params := FilterParams(c)
	QParams := PMap(params)
	var conds []gen.Condition

	for param, value := range QParams {
		if condFunc, ok := paramToCondition[param]; ok {
			conds = append(conds, condFunc(value))
		}
	}
	return conds
}

func FilterParams(c *fiber.Ctx) map[string]string {
	params := c.Queries()
	filteredParams := make(map[string]string)
	for key, value := range params {
		if value != "" {
			lowerKey := strings.ToLower(key)
			lowerValue := strings.ToLower(value)
			filteredParams[lowerKey] = lowerValue
		}
	}
	return filteredParams
}

func PMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = "%" + v + "%"
	}
	return result
}
