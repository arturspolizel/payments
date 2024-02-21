package model

import (
	"errors"

	"github.com/arturspolizel/payments/utils"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {

	// Migrate the schema
	err := db.AutoMigrate(&User{}, &ValidationEmail{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not migrate database")
	}

	return &UserRepository{
		database: db,
	}
}

func (r UserRepository) Get(id uint) (User, error) {
	user := User{}
	result := r.database.First(&user, id)

	if result.Error != nil {
		//log
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, &utils.ErrDatabaseNotFound{
				EntityType: "user",
				EntityId:   id,
			}
		}
		return user, result.Error
	}

	return user, nil
}

func (r UserRepository) GetByEmail(email string) (User, error) {
	user := User{}
	result := r.database.Where("email = ?", email).First(&user)

	if result.Error != nil {
		//log
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, &utils.ErrDatabaseNotFound{
				EntityType: "user",
				EntityId:   0,
			}
		}
		return user, result.Error
	}

	return user, nil
}

func (r UserRepository) Update(user User) error {
	result := r.database.Save(&user)
	return result.Error
}

func (r UserRepository) Create(user User) (uint, error) {
	result := r.database.Create(&user)
	return user.ID, result.Error
}

func (r UserRepository) CreateValidationEmail(email ValidationEmail) (uint, error) {
	result := r.database.Create(&email)
	return email.ID, result.Error
}

func (r UserRepository) GetEmailByCode(code string) (ValidationEmail, error) {
	email := ValidationEmail{}
	result := r.database.Where("code = ?", email).First(&email)

	if result.Error != nil {
		//log
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return email, &utils.ErrDatabaseNotFound{
				EntityType: "validationEmail",
				EntityId:   0,
			}
		}
		return email, result.Error
	}

	return email, nil
}

func (r UserRepository) UpdateValidationEmail(email ValidationEmail) error {
	result := r.database.Save(&email)
	return result.Error
}
