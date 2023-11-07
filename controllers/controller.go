package controllers

import (
	"net/http"

	"github.com/1Nelsonel/Company/database"
	"github.com/1Nelsonel/Company/model"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func GetCompanies(ctx *gin.Context)  {
	db := database.DBConn

	var companies []model.Company

	if err := db.Find(&companies).Error; err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Unable to find a Company"})
		return
	}

	// Log a success message in the terminal and display the count of companies
	success := color.New(color.FgHiGreen)
 	success.Printf("Fetch successfully. Total companies: %d\n", len(companies))


	// Create a Gin context
	context := make(map[string]interface{})
	context["companies"]=companies
	// ctx.HTML(http.StatusOK, "index.html", context)
	ctx.JSON(http.StatusOK, companies)
}

// Create company
func CreateCompany(ctx *gin.Context)  {
	db := database.DBConn

	// var company model.Company
	var companies []model.Company

	if err := ctx.ShouldBindJSON(&companies); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 // Create each company individually
	for _, company := range companies {
        if err := db.Create(&company).Error; err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Company"})
            return
        }
        // Print success message for each successfully created company
        green := color.New(color.FgGreen)
    	green.Println("Company created successfully.......")
    }

	ctx.JSON(http.StatusCreated, companies)
}

// UpdateCompany updates an existing Company.
func UpdateCompany(c *gin.Context) {
    var company model.Company
    companyID := c.Param("id")

    db := database.DBConn

    if err := db.Where("id = ?", companyID).First(&company).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "company not found"})
        return
    }

    if err := c.ShouldBindJSON(&company); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Use db.Save to update the company
    if err := db.Save(&company).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update company"})
        return
    }

	// Log a success message in the terminal
	success := color.New(color.FgGreen)
    success.Printf("Company with ID: %d and Name: %s updated successfully\n", company.ID, company.Name)

    c.JSON(http.StatusOK, company)
}


// DeleteCompany deletes an existing Company.
func DeleteCompany(c *gin.Context) {
	db := database.DBConn

    companyID := c.Param("id")

    if err := db.Where("id = ?", companyID).Delete(&model.Company{}).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}