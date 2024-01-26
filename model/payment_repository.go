package model

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func NewPaymentRepository(host, user, password, databaseName, port string) *PaymentRepository {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Payment{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not migrate database")
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
