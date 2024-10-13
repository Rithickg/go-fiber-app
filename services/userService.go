package services

import (
	"my-fiber-app/database"
	"my-fiber-app/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func DeleteUser(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}
