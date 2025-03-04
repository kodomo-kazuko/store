package main

import (
	"store/database"
	"store/models"

	"gorm.io/gen"
	// or the appropriate driver for your database
)

type Querier interface {
	// SELECT * FROM @@table WHERE active = true
	IsActive() ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: false,
	})

	// Use the connected database
	g.UseDB(database.Database.DB)

	// Apply models
	g.ApplyBasic(models.Models...)

	// Execute the generator
	g.Execute()
}
