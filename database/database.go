package database

import (
	"github.com/1Nelsonel/Company/model"
	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	var err error

	db, err := gorm.Open(sqlite.Open("company.db"), &gorm.Config{})
	if err != nil {
		danger := color.New(color.FgHiRed)
		danger.Println("Error connecting to the database.....")
	}
	green := color.New(color.FgGreen)
    green.Println("Database connected successfully.......")

	// Auto Migration
	if err := db.AutoMigrate(&model.Company{}); err != nil {
        danger := color.New(color.FgHiRed)
        danger.Println("Auto migration failed:: ", err)
    } else {
        cyan := color.New(color.FgHiCyan)
        cyan.Println("Auto migration successful......")
    }

	DBConn = db

}