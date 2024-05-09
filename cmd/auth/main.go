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

	// var key = []byte("-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAOpQ9mFP3TcwIzYfAt4DBoOfFyaXAi59ti2rFe4umtNA=\n-----END PUBLIC KEY-----")
	// var signingKey = []byte("-----BEGIN PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIM0RLoe/ASJtOWt3QUZ0bd6J0rGMI4m/LrKxueAL95AV\n-----END PRIVATE KEY-----")

	keyEnv := os.Getenv("JWT_PUBLICKEY")
	key := []byte(keyEnv)
	signKeyEnv := os.Getenv("JWT_PRIVATEKEY")
	signingKey := []byte(signKeyEnv)

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}

	userRepo := model.NewUserRepository(db)
	emailAdapter := controller.NewEmailAdapter()
	jwtProcessor := utils.NewJwtProcessorWithPrivate(key, signingKey, jwt.SigningMethodEdDSA)
	userController := controller.NewUserController(userRepo, emailAdapter, jwtProcessor)

	engine := gin.Default()
	router := engine.Group("/auth")

	userHandler := handler.NewUserHandler(userController, router)
	userHandler.SetRouters()

	engine.Run(":8080")
}
