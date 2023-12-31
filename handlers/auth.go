package handlers

import (
	"errors"
	"log"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/aminkhn/mysql-rest-api/config"
	"github.com/aminkhn/mysql-rest-api/database"
	"github.com/aminkhn/mysql-rest-api/logic"
	"github.com/aminkhn/mysql-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// Finds user by Email and returns User
func getUserByEmail(e string) (*models.User, error) {
	db := database.Database.Db
	var user models.User
	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Finds user by Username and returns User
func getUserByUsername(u string) (*models.User, error) {
	db := database.Database.Db
	var user models.User
	if err := db.Where(&models.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Checks if input is Email or not
func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login gets user and password and gives JWT token
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	identity := input.Identity
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	if isEmail(identity) {
		userModel, err = getUserByEmail(identity)
	} else {
		userModel, err = getUserByUsername(identity)
	}

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	if !logic.CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// loading Env variables
	loadConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("can not load Envirnment variables", err)
	}
	t, err := token.SignedString([]byte(loadConfig.Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

// logout adds JWT token in blacklist
func Logout(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) == 2 {
		reqToken = splitToken[1]
		claims := jwt.MapClaims{}
		// loading Env variables
		loadConfig, err := config.LoadConfig("./")
		if err != nil {
			log.Fatal("can not load Envirnment variables", err)
		}
		token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(loadConfig.Secret), nil
		})
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
		}
		userId := strconv.FormatFloat(claims["user_id"].(float64), 'f', -1, 64)

		_, err = database.RedisDb.Db.Set(userId, token.Raw, time.Hour*1).Result()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "token is missing"})
}
