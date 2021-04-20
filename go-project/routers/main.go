package routes

import (
	"github.com/achmadsy/go-project/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router = gin.Default()

//Run the gin route
func Run() {
	getRoutes()
	setSwagger()
	Router.Run(":8080")
}

func getRoutes() {
	v1 := Router.Group("/v1")
	addUserRoutes(v1)
}

func setSwagger() {
	docs.SwaggerInfo.Title = "User's API"
	docs.SwaggerInfo.Description = "This is a user's API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
