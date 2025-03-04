package main

import (
	"context"
	"log"
	"store/config"
	"store/database"
	"store/models"
	"store/query"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.MustLoad()

	// Connect to the database
	database.ConnectDb()

	// Initialize query
	query.SetDefault(database.Database.DB)

	// Create a context
	ctx := context.Background()

	q := query.Use(database.Database.DB)

	position := &models.Position{
		Name: "position_test",
	}

	// Seed the database
	if err := q.Position.WithContext(ctx).Create(position); err != nil {
		log.Println(err)
	}

	roles := []*models.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	// Seed the database
	if err := q.Role.WithContext(ctx).CreateInBatches(roles, len(roles)); err != nil {
		log.Println(err)
	}

	organization := &models.Organization{
		Name:             "test_company",
		Email:            "test@gmail.com",
		Phone:            "12345678",
		Register:         "12345678",
		DetailedLocation: "test",
	}

	// Seed the database
	if err := q.Organization.WithContext(ctx).Create(organization); err != nil {
		log.Println(err)
	}

	user := &models.User{
		FirstName:      "first_name_test",
		LastName:       "last_name_test",
		Email:          "s123@gmail.com",
		Password:       "Admin@123",
		Phone:          "12345678",
		RoleID:         1,
		PositionID:     1,
		OrganizationID: 1,
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	user.Password = string(hashedPassword)

	// Seed the database
	if err := q.User.WithContext(ctx).Create(user); err != nil {
		log.Println(err)
	}
}
