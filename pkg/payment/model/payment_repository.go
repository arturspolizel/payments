package model

import (
	"errors"
	"time"

	"github.com/arturspolizel/payments/utils"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	database *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {

	// Migrate the schema
	err := db.AutoMigrate(&Payment{}, &Refund{})

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
			return payment, &utils.ErrDatabaseNotFound{
				EntityType: "payment",
				EntityId:   id,
			}
		}
		return payment, result.Error
	}

	return payment, nil
}

func (r *PaymentRepository) List(startId, pageSize uint, startDate, endDate time.Time) ([]Payment, error) {
	var payments []Payment
	// TODO: early loading?
	result := r.database.Model(&Payment{}).
		Preload("Merchant").
		Preload("Refunds").
		Scopes(Paginate(startId, pageSize)).
		Where("created_at between ? and ?", startDate, endDate).
		Find(&payments)

	if result.Error != nil {
		// log error
		return payments, result.Error
	}

	return payments, nil
}

func (r *PaymentRepository) Create(payment Payment) (uint, error) {
	result := r.database.Create(&payment)
	return payment.ID, result.Error
}

func (r *PaymentRepository) CreateRefund(refund Refund) (uint, error) {
	result := r.database.Create(&refund)
	return refund.ID, result.Error
}

func (r *PaymentRepository) Update(payment Payment) error {
	result := r.database.Save(&payment)
	return result.Error
}
