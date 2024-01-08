package main

import (
	"github.com/AndreCordeir0/fullstack-challenge/backend/cmd/order"
	"github.com/AndreCordeir0/fullstack-challenge/backend/cmd/pizza"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/pizzas", pizza.Create)
	r.GET("/api/pizzas", pizza.Get)
	r.GET("/api/orders", order.Get)
	r.GET("/api/orders/{id}", order.GetById)
	r.POST("/api/orders", order.Create)

	r.Run(":8080")
}
