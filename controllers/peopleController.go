package controllers

import (
	"gorm-pagination/initializers"
	"gorm-pagination/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
}

func PeopleIndexGet(c *gin.Context) {

	//Get page number
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	//calculate total number pages
	var totalRows int64
	initializers.DB.Model(&models.Person{}).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

	//calculate offset parameter
	offset := (page - 1) * perPage

	//Get the people index
	var people []models.Person
	// db.Find(&people)
	initializers.DB.Limit(perPage).Offset(offset).Find(&people)

	//Render the page
	c.HTML(http.StatusOK, "index.html", gin.H{
		"people": people,
		"pagination": PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
			TotalPages:   int(totalPages),
			TwoAfter:     page + 2,
			TwoBelow:     page - 2,
		},
	})

	// c.HTML(http.StatusOK, "index2.html", nil)
}

// func testIndex(c *gin.Context) {
// 	c.HTML(http.StatusOK,"index2.html", nil)
// }
