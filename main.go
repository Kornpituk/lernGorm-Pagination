package main

import (
	"gorm-pagination/initializers"
	"gorm-pagination/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name  string
	Email uint
}

func init() {
	initializers.ConnectToDB()
	initializers.SyncDB()
	// initializers.CreatePeople()
}

func main() {

	//create gin app
	app := gin.Default()

	//Cofigure app
	app.LoadHTMLGlob("templates/**/*")
	app.Static("assets/", "./assets")

	//Routing
	app.GET("/",controllers.PeopleIndexGet)
	app.GET("/page/:page", controllers.PeopleIndexGet)
	// app.GET("/index",controllers.PeopleIndexGet)

	app.Run()

}
