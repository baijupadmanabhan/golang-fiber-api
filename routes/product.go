package routes

import (
	"errors"

	"github.com/baijupadmanabhan/golang-fiber-api/database"
	"github.com/baijupadmanabhan/golang-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func createResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := createResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)
	responseProduct := []Product{}

	for _, product := range products {
		responseProduct = append(responseProduct, createResponseProduct(product))
	}

	return c.Status(200).JSON(responseProduct)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("Product does not exists")
	}

	return nil
}

func GetProductById(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that the :id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := createResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that the :id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData updateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := createResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that the :id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted Product")

}
