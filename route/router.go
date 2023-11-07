package route

import (
	"github.com/1Nelsonel/Company/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine)  {
	r.GET("/", controllers.GetCompanies)
	r.POST("company/create/", controllers.CreateCompany)
	r.PUT("company/:id/update/", controllers.UpdateCompany)
	r.DELETE("company/:id/delete/", controllers.DeleteCompany)
}