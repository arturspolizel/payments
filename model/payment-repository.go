package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func New(host, user, password, databaseName string) *PaymentRepository {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		//log error, fail
	}

	return &PaymentRepository{
		database: db,
	}
}
