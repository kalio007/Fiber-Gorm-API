package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kalio007/Fiber-Gorm-API/database"
	"github.com/kalio007/Fiber-Gorm-API/models"
)

type User struct {
	//not a model but a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt(":id")
	var user models.User
	if err != nil {
		return c.Status(200).JSON("Please ensure the that the :id i an integer")
	}
	if err := findUser(id, &user); err != nil {
		c.Status(200).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
