package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/shbhom/fiber-api/database"
	"github.com/shbhom/fiber-api/models"
)

// not a user Model, this is a serializer
type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func createResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Email: userModel.Email, Name: userModel.Name}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&user)
	responseUser := createResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		responseUser := createResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.DB.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user doesn't exists")
	}
	return nil
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please make sure the id is an Uint")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseUser := createResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is Uint")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type updatedData struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var updatedUser updatedData

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	database.DB.Db.Save(&user)

	// database.DB.Db.Model(&user).Updates(User{Name: updatedUser.Name, Email: updatedUser.Email})
	responseUser := createResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is Uint")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseUser := createResponseUser(user)
	database.DB.Db.Delete(&user)
	return c.Status(200).JSON(responseUser)
}
