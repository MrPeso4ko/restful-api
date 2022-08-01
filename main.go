package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Diploma struct {
	Id      int    `json:"id"`
	Year    int    `json:"year"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var diplomas = []Diploma{{1, 2022, "Petr", "Salnikov"}}
var diplomasId = map[int]*Diploma{1: &diplomas[0]}
var diplomasYear = map[int][]*Diploma{2022: {&diplomas[0]}}

func getDiplomas(c *gin.Context) {
	year, err := strconv.Atoi(c.DefaultQuery("year", "-1"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if year == -1 {
		c.IndentedJSON(http.StatusOK, diplomas)
	} else {
		if diplomasYear[year] != nil {
			c.IndentedJSON(http.StatusOK, diplomasYear[year])
		} else {
			c.IndentedJSON(http.StatusOK, []Diploma{})
		}
	}
}

func addDiplomas(c *gin.Context) {
	var newDiplomas []Diploma
	err := c.BindJSON(&newDiplomas)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	n := len(diplomas)
	diplomas = append(diplomas, newDiplomas...)
	for i := 0; i < len(newDiplomas); i++ {
		diplomasId[newDiplomas[i].Id] = &diplomas[n+i]
		diplomasYear[newDiplomas[i].Year] = append(diplomasYear[newDiplomas[i].Year], &diplomas[n+i])
	}
	c.IndentedJSON(http.StatusCreated, newDiplomas)
}

func removeDiploma(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	diploma := diplomasId[id]
	if diploma == nil {
		c.Data(http.StatusNotFound, "text/plain", []byte("Diploma not found"))
		return
	}
	for i := 0; i < len(diplomas); i++ {
		if diplomas[i].Id == id {
			diplomas = append(diplomas[:i], diplomas[i+1:]...)
			break
		}
	}
	for i := 0; i < len(diplomasYear[diploma.Year]); i++ {
		if diplomasYear[diploma.Year][i].Id == id {
			diplomasYear[diploma.Year] = append(diplomasYear[diploma.Year][:i], diplomasYear[diploma.Year][i+1:]...)
		}
	}
	delete(diplomasId, id)
	c.Data(http.StatusOK, "text/plain", []byte("Diploma deleted"))
}

func main() {
	router := gin.Default()
	router.GET("/diploma/list", getDiplomas)
	router.POST("/diploma/list", addDiplomas)
	router.DELETE("/diploma/list", removeDiploma)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
