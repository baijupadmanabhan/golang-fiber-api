package routes

import (
	"errors"

	"github.com/baijupadmanabhan/golang-fiber-api/database"
	"github.com/baijupadmanabhan/golang-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)
	responseUser := []User{}

	for _, user := range users {
		responseUser = append(responseUser, createResponseUser(user))
	}

	return c.Status(200).JSON(responseUser)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user does not exists")
	}

	return nil
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that the :id is integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := createResponseUser(user)

	return c.Status(200).JSON(responseUser)

}
