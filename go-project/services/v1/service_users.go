package servicecharsv1

import (
	"github.com/achmadsy/go-project/db"
	"github.com/achmadsy/go-project/models"
)

//GetAllUsers to query to get all user characters
func GetAllUsers(users *[]models.User) error {
	if err := db.DB.Preload("Addressses").Find(&users).Error; err != nil {
		return err
	}
	return nil
}

//UpdateDetailUser to update user character
func UpdateDetailUser(userID *uint64, userData *models.User) error {
	if err := db.DB.Model(&userData).Where("id = ?", userID).Updates(models.User{Name: userData.Name}).Error; err != nil {
		return err
	}
	return nil
}

//PostNewChar to post new user character
func PostNewUser(userData *models.User) error {
	if result := db.DB.Create(&userData); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUserByID to get user by ID
func GetUserByID(userID *uint64) (models.User, error) {
	user := models.User{}
	if err := db.DB.Model(&user).Where("id = ?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
