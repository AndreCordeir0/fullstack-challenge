package order

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/AndreCordeir0/fullstack-challenge/backend/cmd/pizza"
	"github.com/AndreCordeir0/fullstack-challenge/backend/database"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Adress     string      `json:"adress"`
	OrderItems []OrderItem `json:"order_items"`
}
type OrderItem struct {
	Id       int
	IdPizza  int         `json:"id_pizza"`
	Quantity int         `json:"quantity"`
	IdOrder  int         `json:"-"`
	Pizza    pizza.Pizza `json:"-"`
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

	orderId, err := order.CreateOrder(tx)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err})
		return
	}

	orderItemsError := order.CreateOrderItems(tx, orderId)
	if orderItemsError != nil {
		fmt.Println(orderItemsError)
		c.JSON(500, gin.H{"error": orderItemsError})
		return
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		fmt.Println("Commit error", errCommit)
		c.JSON(500, errCommit)
		return
	}

	location := fmt.Sprintf("orders/%d", orderId)
	c.Writer.Header().Add("Location", location)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "created order",
	})
}

func Get(c *gin.Context) {
	db := database.GetConnection()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Error getting transaction", err)
	}
	defer db.Close()

	result, queryErr := tx.Query("SELECT * FROM ORDER_PIZZA")
	if queryErr != nil {
		c.AbortWithStatusJSON(500, queryErr.Error())
	}

	var orders []Order
	for result.Next() {
		var order Order
		if err := result.Scan(&order.Id, &order.Name, &order.Adress); err != nil {
			c.JSON(500, err.Error())
			return
		}
		orderItems, err := GetOrderItems(order.Id)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		order.OrderItems = orderItems
		orders = append(orders, order)
	}

	c.JSON(200, orders)
}

func GetOrderItems(orderId int) ([]OrderItem, error) {
	db := database.GetConnection()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Error getting transaction", err)
	}
	defer db.Close()

	result, queryErr := tx.Query("SELECT * FROM ORDER_ITEM o WHERE o.id_order = $1", orderId)
	if queryErr != nil {
		return []OrderItem{}, queryErr
	}
	var orderItems []OrderItem
	for result.Next() {
		var orderItem OrderItem
		if err := result.Scan(&orderItem.Id, &orderItem.IdPizza, &orderItem.Quantity, &orderItem.IdOrder); err != nil {
			return []OrderItem{}, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

func (o *Order) CreateOrder(tx *sql.Tx) (lastInsertedId int, txError error) {
	txError = tx.QueryRow("INSERT INTO ORDER_PIZZA (name, adress) VALUES ($1, $2) RETURNING id", o.Name, o.Adress).Scan(&lastInsertedId)
	if txError != nil {
		return 0, txError
	}
	return lastInsertedId, nil
}

func (o *Order) CreateOrderItems(tx *sql.Tx, orderId int) error {
	fmt.Println("Order id: ", orderId)
	for _, item := range o.OrderItems {
		_, txError := tx.Exec("INSERT INTO ORDER_ITEM (id_pizza, quantity, id_order) VALUES ($1, $2, $3)", item.IdPizza, item.Quantity, orderId)
		if txError != nil {
			return txError
		}
	}
	return nil
}

func GetById(c *gin.Context) {

}
