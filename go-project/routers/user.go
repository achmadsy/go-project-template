package routes

import (
	apiusers_v1 "github.com/achmadsy/go-project/api/v1"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user/")

	users.GET("/all", apiusers_v1.GetAllUser)
	users.POST("/", apiusers_v1.PostNewUser)
	users.PUT("/:id", apiusers_v1.PutEditedUser)
}
