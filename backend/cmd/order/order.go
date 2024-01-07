package order

import (
	"fmt"
	"log"

	"github.com/AndreCordeir0/fullstack-challenge/backend/cmd/pizza"
	"github.com/AndreCordeir0/fullstack-challenge/backend/database"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Pizza    pizza.Pizza
	IdPizza  int `json:"id_pizza"`
	Quantity int `json:"quantity"`
}

func Create(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := database.GetConnection()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Error getting transaction")
	}
	defer db.Close()
	defer tx.Rollback()
	result, txError := tx.Exec("INSERT INTO ORDER_PIZZA (id_pizza, quantity) VALUES (?, ?)", order.IdPizza, order.Quantity)
	if txError != nil {
		//TODO
		c.JSON(500, "")
		return
	}
	id, idError := result.LastInsertId()
	if idError != nil {
		//TODO

		c.JSON(500, "")
		return
	}
	location := fmt.Sprintf("orders/%d", id)
	c.Writer.Header().Add("Location", location)
}

func Get(c *gin.Context) {

}

func GetById(c *gin.Context) {

}
