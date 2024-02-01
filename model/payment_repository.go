package model

import (
	"errors"

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
	result := r.database.First(&payment, id)

	if result.Error != nil {
		//log
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return payment, &ErrDatabaseNotFound{
				entityType: "payment",
				entityId:   id,
			}
		}
		return payment, result.Error
	}

	return payment, nil
}

func (r *PaymentRepository) Create(payment Payment) (uint, error) {
	result := r.database.Create(&payment)
	return payment.ID, result.Error
}
