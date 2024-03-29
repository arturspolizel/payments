package main

import (
	"fmt"
	"os"

	"github.com/arturspolizel/payments/pkg/payment/controller"
	"github.com/arturspolizel/payments/pkg/payment/handler"
	"github.com/arturspolizel/payments/pkg/payment/model"
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
	var key = []byte("-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAOpQ9mFP3TcwIzYfAt4DBoOfFyaXAi59ti2rFe4umtNA=\n-----END PUBLIC KEY-----")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", "123", "payments", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}

	paymentRepo := model.NewPaymentRepository(db)
	merchantRepo := model.NewMerchantRepository(db)
	paymentController := controller.NewPaymentController(paymentRepo)
	merchantController := controller.NewMerchantController(merchantRepo)

	jwtProcessor := utils.NewJwtProcessor(key, jwt.SigningMethodEdDSA)

	engine := gin.Default()
	engine.Use(utils.JwtMiddleware(jwtProcessor))
	router := engine.Group("/payment")

	paymentHandler := handler.NewPaymentHandler(paymentController, router)
	paymentHandler.SetRouters()

	merchantHandler := handler.NewMerchantHandler(merchantController, router)
	merchantHandler.SetRouters()

	engine.Run(":8080")
}
