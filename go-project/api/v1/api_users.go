package users

import (
	"net/http"
	"strconv"

	"github.com/achmadsy/go-project-template/go-project/models"
	serviceusers_v1 "github.com/achmadsy/go-project-template/go-project/services/v1"
	"github.com/gin-gonic/gin"
)

//GetAllUsers is to get all users
// @Summary Show all registered users
// @Description get all users
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/all [get]
func GetAllUser(c *gin.Context) {
	users := []models.User{}

	if err := serviceusers_v1.GetAllUsers(&users); err != nil {
		out := models.SetResponse(false, "Failed to fetch users", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}

	c.JSON(http.StatusOK, users)
}

//PostNewUser is to post new user
// @Summary Post new user
// @Description Post new user
// @Produce json
// @Accept  json
// @Param user body models.User true "New User"
// @Success 200 {object} models.User
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/ [post]
func PostNewUser(c *gin.Context) {
	newUser := models.User{}
	c.Bind(&newUser)
	newUser.ID = 0
	if err := serviceusers_v1.PostNewUser(&newUser); err != nil {
		out := models.SetResponse(false, "Failed to create user", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}
	c.JSON(http.StatusOK, newUser)
}

//PutEditedUser is to edit the user data
// @Summary Edit user
// @Description Edit user detail
// @Produce json
// @Accept  json
// @Param user body models.User true "Edit User Data"
// @Param userID path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/{userID}/ [put]
func PutEditedUser(c *gin.Context) {
	user := models.User{}
	userID, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	c.Bind(&user)
	if err := checkUserID(userID); err {
		out := models.SetResponse(false, "UserID not found", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := serviceusers_v1.UpdateDetailUser(&userID, &user); err != nil {
		out := models.SetResponse(false, "Failed to update user detail", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}
	out := models.SetResponse(true, "Update user detail success", nil)
	c.JSON(http.StatusOK, out)
}

func checkUserID(userID uint64) bool {
	if user, _ := serviceusers_v1.GetUserByID(&userID); user.ID != 0 {
		return false
	}
	return true
}
