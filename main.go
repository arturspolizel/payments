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

	engine := gin.Default()
	router := engine.Group("/payment")

	paymentHandler := handler.NewPaymentHandler(paymentController, router)
	paymentHandler.SetRouters()

	engine.Run(":8080")
}
