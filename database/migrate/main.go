package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"store/config"
	"store/database"
	"store/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	config.MustLoad()
	cfg := config.Get()

	// Set up context and connection string
	ctx := context.Background()
	connStr := database.GenerateDSN(*cfg.DB)

	// Connect to database using pgxpool
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// Connect to database using GORM
	database.ConnectDb()
	db := database.Database.DB

	// Run migrations
	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}
}

func Migrate(db *gorm.DB) error {

	// Drop tables
	if err := db.Migrator().DropTable(models.Models...); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	// Auto-migrate the models
	return db.AutoMigrate(models.Models...)
}
