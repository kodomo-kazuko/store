package middleware

import (
	"fmt"
	"net/url"
	"os"
	"store/form"
	"store/shared"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		jwtSecret = []byte(secret)
	} else {
		jwtSecret = []byte("hello mr blue sky")
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(claims map[string]interface{}) (string, error) {
	jwtClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expiration (1 day)
	}
	for key, value := range claims {
		jwtClaims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(jwtSecret)
}

// JWTMiddleware returns the Fiber JWT middleware
func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: jwtSecret},
		// Pass jwtSecret as a byte slice
		ContextKey:   "user",          // Key to store user information in context
		ErrorHandler: jwtErrorHandler, // Custom error handler
	})
}

// jwtErrorHandler handles JWT authentication errors
func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(form.Response{
		Success: false,
		Message: "Unauthorized: " + err.Error(),
	})
}

// GetJWTToken returns the JWT token from the context
func GetJWTToken(c *fiber.Ctx) (string, error) {
	user := c.Locals("user").(*jwt.Token)
	return user.Raw, nil
}

// ExtractClaimFromToken retrieves a specific claim from the JWT token
func ExtractClaimFromToken(c *fiber.Ctx, claimKey string) (interface{}, error) {
	// Retrieve the token from the context
	tokenObj, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("unable to extract token from context")
	}

	// Extract claims
	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to extract claims")
	}

	// Retrieve the specific claim
	value, exists := claims[claimKey]
	if !exists {
		return nil, fmt.Errorf("claim '%s' not found in token", claimKey)
	}

	return value, nil
}

func ExtractOrganizationID(c *fiber.Ctx) uint {
	// Extract the claim "organization_id" from the token
	orgId, err := ExtractClaimFromToken(c, "organization_id")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(form.Response{
			Success: false,
			Error:   err.Error(),
		})
		return 0
	}

	// Ensure orgId is a float64 before conversion
	orgIdFloat, ok := orgId.(float64)
	if !ok {
		c.Status(fiber.StatusInternalServerError).JSON(form.Response{
			Success: false,
			Error:   "invalid organization ID format",
		})
		return 0
	}

	// Convert float64 to uint
	orgIdUint := uint(orgIdFloat)
	return orgIdUint

}

func ExtractUserID(c *fiber.Ctx) uint {
	// Extract the claim "user_id" from the token
	userId, err := ExtractClaimFromToken(c, "user_id")
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(form.Response{
			Success: false,
			Error:   err.Error(),
		})
		return 0
	}

	// Ensure userId is a float64 before conversion
	userIdFloat, ok := userId.(float64)
	if !ok {
		c.Status(fiber.StatusInternalServerError).JSON(form.Response{
			Success: false,
			Error:   "invalid user ID format",
		})
		return 0
	}

	// Convert float64 to uint
	userIdUint := uint(userIdFloat)
	return userIdUint
}

func ExtractAdmin(c *fiber.Ctx) bool {
	// Extract the claim "permissions" from the token
	permissions, err := ExtractClaimFromToken(c, "permissions")
	if err != nil {
		shared.InternalServerError(c, err)
		return false
	}

	// Check if permissions is a string and if it is "*"
	if permStr, ok := permissions.(string); ok && permStr == "*" {
		return true
	}

	return false
}

func GetIDFromParams(c *fiber.Ctx) uint {
	id, _ := c.ParamsInt("id")

	return uint(id)
}

func GetStringFromParams(c *fiber.Ctx) (string, error) {
	// Get the "id" parameter from the URL
	param := c.Params("id")

	// Decode the URL-encoded string
	decodedParam, err := url.QueryUnescape(param)
	if err != nil {
		// If decoding fails, fallback to the original param
		decodedParam = param
	}

	// Try to convert the decoded param to an integer
	if _, err := strconv.Atoi(decodedParam); err != nil {
		// If conversion fails, it means the param is a non-numeric string, so return it
		return decodedParam, nil
	}

	// If conversion succeeds (i.e. it's numeric), return an empty string
	return decodedParam, err
}
