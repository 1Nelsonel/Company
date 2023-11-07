package main

import (
	"github.com/1Nelsonel/Company/database"
	"github.com/1Nelsonel/Company/route"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main()  {

	// Database connect
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connection")
	}

	defer sqlDb.Close()
	

	app := gin.Default()

	// Templates
	app.LoadHTMLGlob("templates/*")

	// Initialize routes by calling SetupRoutes
	route.SetupRoutes(app)

	app.Run()
	
}