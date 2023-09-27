package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/shbhom/fiber-api/database"
	"github.com/shbhom/fiber-api/models"
)

type Product struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}

func createResponseProduct(prod models.Product) Product {
	return Product{ID: prod.ID, Title: prod.Title, Price: prod.Price}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	database.DB.Db.Create(&product)
	responseProduct := createResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.DB.Db.Find(&products)

	responseProducts := []Product{}

	for _, prod := range products {
		responseProduct := createResponseProduct(prod)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.DB.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("no product found")
	}
	return nil
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var prod models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findProduct(id, &prod); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseProduct := createResponseProduct(prod)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type updatedData struct {
		Title string `json:"title"`
		Price uint   `json:"price"`
	}

	var updatedProduct updatedData

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Title = updatedProduct.Title
	product.Price = updatedProduct.Price
	database.DB.Db.Save(product)

	// database.DB.Db.Model(&product).Updates(models.Product{Title: updatedProduct.Title, Price: updatedProduct.Price})

	responseProduct := createResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseProduct := createResponseProduct(product)
	database.DB.Db.Delete(&product)
	return c.Status(200).JSON(responseProduct)
}
