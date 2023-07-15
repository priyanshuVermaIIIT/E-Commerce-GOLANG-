package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyanshuVermaIIIT/Ecommerce/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/user/signup", controllers.SignUp())
	incommingRoutes.POST("/user/login", controllers.Login())
	incommingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incommingRoutes.GET("/user/productview", controllers.SearchProduct())
	incommingRoutes.GET("/user/search", controllers.SearchProductByQuery())

}
