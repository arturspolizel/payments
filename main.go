package main

import (
	"fmt"

	"github.com/arturspolizel/payments/controller"
	"github.com/arturspolizel/payments/handler"
	"github.com/arturspolizel/payments/model"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running server")

	paymentRepo := model.NewPaymentRepository("localhost", "postgres", "123", "payments", "5432")
	paymentController := controller.NewPaymentController(paymentRepo)

	router := gin.Default()

	paymentHandler := handler.NewPaymentHandler(paymentController, router)
	paymentHandler.SetRouters()

	router.Run(":8080")
}
