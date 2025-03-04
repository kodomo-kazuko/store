package database

import (
	"fmt"
	"log"
	"os"
	"store/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	config := config.Get()
	dsn := GenerateDSN(*config.DB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(2)
	}

	// Initialize GORM Logger
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DbInstance{DB: db}
	log.Println("connected to database")
}

func GenerateDSN(dbConfig config.DatabaseConfig) string {
	return fmt.Sprintf(
		`host=%s port=%d dbname=%s user=%s password=%s search_path=%s sslmode=disable TimeZone=Asia/Ulaanbaatar`,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Schema,
	)
}
