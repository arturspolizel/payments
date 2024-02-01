package model

import (
	"errors"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	database *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {

	// Migrate the schema
	err := db.AutoMigrate(&Merchant{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not migrate database")
	}

	return &MerchantRepository{
		database: db,
	}
}

func (r *MerchantRepository) Get(id uint) (Merchant, error) {
	merchant := Merchant{}
	result := r.database.First(&merchant, id)

	if result.Error != nil {
		//log
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return merchant, &ErrDatabaseNotFound{
				entityType: "merchant",
				entityId:   id,
			}
		}
		return merchant, result.Error
	}

	return merchant, nil
}

func (r *MerchantRepository) Create(merchant Merchant) (uint, error) {
	result := r.database.Create(&merchant)
	return merchant.ID, result.Error
}
