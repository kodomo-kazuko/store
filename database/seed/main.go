package main

import (
	"context"
	"log"
	"store/config"
	"store/database"
	"store/models"
	"store/query"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	config.MustLoad()

	// Connect to the database
	database.ConnectDb()

	// Initialize query (optional, if needed for non-transactional operations)
	query.SetDefault(database.Database.DB)

	// Create a context
	ctx := context.Background()

	// Wrap seed operations in a transaction
	err := database.Database.DB.Transaction(func(tx *gorm.DB) error {
		// Use the transaction for the query builder
		q := query.Use(tx)

		// Seed roles
		roles := []*models.Role{
			{Name: "admin"},
			{Name: "user"},
		}
		if err := q.Role.WithContext(ctx).CreateInBatches(roles, len(roles)); err != nil {
			return err
		}

		// Seed organization
		organization := &models.Organization{
			Name:     "test_company",
			Email:    "test@gmail.com",
			Phone:    "12345678",
			Register: "12345678",
			Address:  "test address",
		}
		if err := q.Organization.WithContext(ctx).Create(organization); err != nil {
			return err
		}

		// Prepare user data
		user := &models.User{
			FirstName:      "first_name_test",
			LastName:       "last_name_test",
			Email:          "s123@gmail.com",
			Password:       "Admin@123", // Plain text, will be hashed next
			Phone:          "12345678",
			RoleID:         1,
			OrganizationID: 1,
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		// Seed user
		if err := q.User.WithContext(ctx).Create(user); err != nil {
			return err
		}

		// If all operations succeed, the transaction is committed.
		return nil
	})

	if err != nil {
		log.Println("Seeding failed:", err)
	} else {
		log.Println("Seeding completed successfully.")
	}
}
