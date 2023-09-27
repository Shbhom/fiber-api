package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shbhom/fiber-api/database"
	"github.com/shbhom/fiber-api/models"
)

type order struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"products"`
	// Amount    uint      `json:"amt"`
	CreatedAt time.Time `json:"orderDate&Time"`
}

func createResponseOrder(ordermodel models.Order, user User, products Product) order {
	return order{ID: ordermodel.ID, User: user, Product: products, CreatedAt: ordermodel.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	var user models.User
	var product models.Product

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := findProduct(order.ProdRefer, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.DB.Db.Create(&order)
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product)

	responseOrder := createResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}

func findOrder(id int, orderModel *models.Order) error {
	database.DB.Db.Find(&orderModel, "id=?", id)
	if orderModel.ID == 0 {
		return errors.New("no order found")
	}
	return nil
}

func GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order
	var user models.User
	var prod models.Product
	if err != nil {
		return c.Status(400).JSON("please make sure :id is an Uint")
	}
	if err := findOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	if err := findProduct(order.ProdRefer, &prod); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	if err := findOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(prod)
	responseOrder := createResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}
