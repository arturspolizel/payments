package main

import (
	"fmt"
	"os"

	"github.com/arturspolizel/payments/pkg/auth/controller"
	"github.com/arturspolizel/payments/pkg/auth/handler"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/arturspolizel/payments/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Running server")

	// TODO: Abstract this out to model package, use env variables
	var key = []byte("testKey")
	var signingKey = []byte("signingKey")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", "123", "payments", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}

	userRepo := model.NewUserRepository(db)
	emailAdapter := controller.NewEmailAdapter()
	jwtProcessor := utils.NewJwtProcessorWithPrivate(key, signingKey, jwt.SigningMethodEdDSA)
	userController := controller.NewUserController(userRepo, emailAdapter, *jwtProcessor)

	engine := gin.Default()
	router := engine.Group("/auth")

	userHandler := handler.NewUserHandler(userController, router)
	userHandler.SetRouters()

	engine.Run(":8080")
}
