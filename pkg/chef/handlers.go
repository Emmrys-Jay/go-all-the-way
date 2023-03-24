package chef

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var chefs []Chef

func init() {
	chefs = make([]Chef, 0)
}

type Chef struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Country    string    `json:"country"`
	YearsOfExp int       `json:"years_of_exp"`
	CreatedAt  time.Time `json:"created_at"`
}

func DeleteChefHandler(c *gin.Context) {
	id := c.Param("chef-id")
	index := -1

	for i := 0; i < len(chefs); i++ {
		if chefs[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "chef not found",
		})
		return
	}

	chefs = append(chefs[:index], chefs[index+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "chef deleted",
	})
}

func UpdateChefHandler(c *gin.Context) {
	id := c.Param("chef-id")

	var chef Chef

	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1

	for i := 0; i < len(chefs); i++ {
		if chefs[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "chef not found",
		})
		return
	}

	chef.Id = id
	chefs[index] = chef

	c.JSON(http.StatusOK, chef)
}

func ListChefsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, chefs)
}

func NewChefHandler(c *gin.Context) {
	var chef Chef

	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chef.Id = xid.New().String()
	chef.CreatedAt = time.Now()
	chefs = append(chefs, chef)
	c.JSON(http.StatusOK, chef)
}
