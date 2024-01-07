package main

import (
	"github.com/AndreCordeir0/fullstack-challenge/internal/order"
	"github.com/AndreCordeir0/fullstack-challenge/internal/pizza"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/pizza", pizza.Create)
	r.GET("/api/pizzas", pizza.Get)
	r.GET("/api/orders", order.Get)
	r.GET("/api/orders/{id}", order.GetById)
}
