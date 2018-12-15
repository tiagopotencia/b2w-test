package main

import (
	"log"
	"net/http"

	"github.com/tiagopotencia/i-am-back/business"

	"github.com/gin-gonic/gin"
)

var DB *business.Database

type response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Content    interface{} `json:"content,omitempty"`
}

func main() {
	db := business.Database{}
	DB = &db
	err := DB.ConnectToDB("mongodb://tp:b2w-test@ds247759.mlab.com:47759/", "b2w-test")

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	group := r.Group("/v1")
	group.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"a": "a",
		})
	})
	group.POST("/planets", AddPlanet)
	group.GET("/planets", GetPlanets)
	group.GET("/planets/:ID", getPlanetByID)
	group.DELETE("/planets/:ID", deletePlanet)

	r.Run()
}

func AddPlanet(c *gin.Context) {
	planet := business.Planet{}

	c.BindJSON(&planet)
	err := business.AddPlanetBusiness(planet, DB)
	err = nil
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, nil)
}

func GetPlanets(c *gin.Context) {

	var result interface{}
	var err error

	planetNameFilter := c.Query("name")

	if planetNameFilter != "" {
		result, err = business.GetPlanetsByName(planetNameFilter, DB)
	} else {
		result, err = business.GetPlanetsBusiness(DB)
	}

	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getPlanetByID(c *gin.Context) {
	ID := c.Param("ID")
	planet, err := business.GetPlanetByID(ID, DB)

	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if planet == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, response{
			StatusCode: http.StatusNotFound,
			Message:    "Planet ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, response{
		StatusCode: http.StatusOK,
		Message:    "Planet found:",
		Content:    &planet,
	})
}

func deletePlanet(c *gin.Context) {
	ID := c.Param("ID")
	planetDeleted, err := business.DeletePlanet(ID, DB)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{
			StatusCode: http.StatusInternalServerError,
			Message:    "An error has occurred. Please try again later",
		})
		return
	}

	if planetDeleted == false {
		c.AbortWithStatusJSON(http.StatusNotFound, response{
			StatusCode: http.StatusNotFound,
			Message:    "Planet ID not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, response{
		StatusCode: http.StatusAccepted,
		Message:    "Planet deleted successfully",
	})
}
