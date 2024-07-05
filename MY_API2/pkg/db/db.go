package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"library2/internal/domain/model"
)

type Database struct {
	db *gorm.DB
}

func Connection() (*Database, error) {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 connect_timeout=10 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	fmt.Println("Successfully connected to database.")

	return &Database{db: db}, nil
}

func (d *Database) GetDb() *gorm.DB {
	return d.db
}

func (d *Database) AutoMigrate(models ...interface{}) error {
	return d.db.AutoMigrate(models...)
}
func (d *Database) Seed() error {
	var count int64
	if err := d.db.Model(model.Book{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {

		books := []model.Book{

			{Title: "Suç ve Ceza", Author: "Dostoyevski", ID: 1},
			{Title: "Nutuk", Author: "G.M.K. Atatürk", ID: 2},
			{Title: "Cehennem", Author: "Dan Brown", ID: 3},
			{Title: "Savaş ve Barış", Author: "Tolstoy", ID: 4},
		}
		for _, book := range books {
			result := d.db.Create(&book)
			if result.Error != nil {
				return fmt.Errorf("error creating book: %v", result.Error)
			}
		}
	}

	return nil
}
