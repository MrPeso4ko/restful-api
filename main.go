package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Diploma struct {
	Id      int    `json:"id"`
	Year    int    `json:"year"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var diplomas = []Diploma{{1, 2022, "Petr", "Salnikov"}}
var diplomas_map = map[int]*Diploma{1: &diplomas[0]}

func get_diplomas(c *gin.Context) { // TODO: optimize
	var diplomas_year []Diploma
	for i := 0; i < len(diplomas); i++ {
		if diplomas[i].Year == 2022 { // TODO: year
			diplomas_year = append(diplomas_year, diplomas[i])
		}
	}
	c.IndentedJSON(http.StatusOK, diplomas_year)
}

func main() {
	router := gin.Default()
	router.GET("/diploma/list", get_diplomas)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
