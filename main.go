package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"` // Hinzufügen des ID-Felds
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
   recipes = make([]Recipe, 0)
   file, _ := os.ReadFile("recipes.json")
   _ = json.Unmarshal([]byte(file), &recipes)
}

func NewRecipeHandler(c *gin.Context) {
   var recipe Recipe

   if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // Schlüssel in Anführungszeichen gesetzt
		})
		return
	}

   recipe.ID = xid.New().String()
   recipe.PublishedAt = time.Now()

   recipes = append(recipes, recipe)

   c.JSON(http.StatusOK, recipe)
}


func ListRecipesHandler(c *gin.Context) {

   c.JSON(http.StatusOK, recipes)

}



func main() {
   router := gin.Default()
   router.POST("/recipes", NewRecipeHandler)
   router.GET("/recipes", ListRecipesHandler)
   router.Run()
}