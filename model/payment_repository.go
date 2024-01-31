package model

import (
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {

	// Migrate the schema
	err := db.AutoMigrate(&Payment{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not migrate database")
	}

	return &PaymentRepository{
		database: db,
	}
}

func (r *PaymentRepository) Get(id uint) (Payment, error) {
	payment := Payment{}
	r.database.Select(&payment, id)

	return payment, nil
}

func (r *PaymentRepository) Create(payment Payment) (uint, error) {
	r.database.Create(&payment)
	return payment.ID, nil
}
