package services

import (
	"my-fiber-app/database"
	"my-fiber-app/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func CreateProduct(product *models.Product) error {
	result := database.DB.Create(product)
	return result.Error
}

func UpdateProduct(product *models.Product) error {
	result := database.DB.Save(product)
	return result.Error
}

func DeleteProduct(id uint) error {
	result := database.DB.Delete(&models.Product{}, id)
	return result.Error
}
