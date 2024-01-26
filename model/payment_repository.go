package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func NewPaymentRepository(host, user, password, databaseName, port string) *PaymentRepository {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Payment{})

	if err != nil {
		//log error, fail
	}

	return &PaymentRepository{
		database: db,
	}
}

func (r *PaymentRepository) Get(id uint) Payment {
	return Payment{}
}

func (r *PaymentRepository) Create(payment Payment) uint {
	r.database.Create(&payment)
	return payment.ID
}
