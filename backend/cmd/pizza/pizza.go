package pizza

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AndreCordeir0/fullstack-challenge/backend/database"
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Price       float64       `json:"price"`
	Ingredients []Ingredients `json:"ingredients"`
}

type Ingredients struct {
	Name string `json:"name"`
}

func Create(c *gin.Context) {
	var pizza Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db := database.GetConnection()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Error getting transaction", err)
	}
	defer db.Close()
	defer tx.Rollback()
	lastInsertId := 0
	j := pizza.IngredientsMarshal()
	txError := tx.QueryRow("INSERT INTO PIZZA (name, price, ingredients) VALUES ($1, $2, $3) RETURNING id", pizza.Name, pizza.Price, j).Scan(&lastInsertId)
	if txError != nil {
		//TODO
		log.Default().Println("Erro inserting pizza", txError)

		c.JSON(500, txError)
		return
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		log.Default().Println("Erro inserting pizza", errCommit)

		c.JSON(500, errCommit)
		return
	}
	location := fmt.Sprintf("pizza/%d", lastInsertId)
	c.Writer.Header().Add("Location", location)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "created",
	})
}
func (p *Pizza) IngredientsMarshal() []byte {
	j, _ := json.Marshal(p.Ingredients)
	return j
}
func Get(c *gin.Context) {
	db := database.GetConnection()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal("Error getting transaction", err)
	}
	defer db.Close()

	result, queryErr := tx.Query("SELECT * FROM PIZZA")
	if queryErr != nil {
		c.AbortWithStatusJSON(500, queryErr.Error())
	}

	var pizzas []Pizza
	for result.Next() {
		var pizza Pizza
		var ingredientsJSON json.RawMessage
		if err := result.Scan(&pizza.Id, &pizza.Name, &pizza.Price, &ingredientsJSON); err != nil {
			c.JSON(500, err.Error())
			log.Fatal(err)
		}
		if err := json.Unmarshal(ingredientsJSON, &pizza.Ingredients); err != nil {
			log.Fatal(err)
		}
		pizzas = append(pizzas, pizza)
	}

	c.JSON(200, pizzas)
}
