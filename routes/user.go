package routes

import (
	"demoapp/database"
	"demoapp/models"

	"github.com/gofiber/fiber/v2"
)

// this is serializer
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Family   string `json:"family"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Family:   userModel.Family,
		Password: userModel.Password,
		Email:    userModel.Email,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
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
