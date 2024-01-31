package model

import (
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
	r.database.Select(&merchant, id)

	return merchant, nil
}

func (r *MerchantRepository) Create(merchant Merchant) (uint, error) {
	r.database.Create(&merchant)
	return merchant.ID, nil
}
