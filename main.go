package main

import (
	"os"

	"github.com/arturspolizel/payments/controller"
	"github.com/arturspolizel/payments/handler"
	"github.com/arturspolizel/payments/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Running server")

	paymentRepo := model.NewPaymentRepository("localhost", "postgres", "123", "payments", "5432")
	paymentController := controller.NewPaymentController(paymentRepo)

	engine := gin.Default()
	router := engine.Group("/payment")

	paymentHandler := handler.NewPaymentHandler(paymentController, router)
	paymentHandler.SetRouters()

	engine.Run(":8080")
}
