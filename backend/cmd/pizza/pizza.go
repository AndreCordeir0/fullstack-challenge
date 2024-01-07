package pizza

import "github.com/gin-gonic/gin"

type Pizza struct {
	Name        string
	Price       string
	Ingredients []Ingredients
}

type Ingredients struct {
	Name string
}

func Create(c *gin.Context) {

}

func Get(c *gin.Context) {

}
