package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/Sparepartku-API/database"
	"github.com/krissukoco/Sparepartku-API/models"
	"golang.org/x/crypto/bcrypt"
)

type userRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type resAuth struct {
	User  models.Person `json:"user"`
	Token string        `json:"token"`
}

func Login(c *fiber.Ctx) error {
	var auth loginUser
	var user models.Person
	if err := c.BodyParser(&auth); err != nil {
		return c.Status(422).JSON(ErrorRes{
			Message:   "Request body cannot be processed",
			ErrorCode: "ERR:BODY:UNPROCESSABLE",
		})
	}
	// Find user by email
	database.DB.First(&user, "email = ?", auth.Email)
	if user.Email == "" {
		return c.Status(403).JSON(ErrorRes{
			Message:   "Email or password is invalid",
			ErrorCode: "ERR:AUTH:INVALID",
		})
	}
	// Validating password
	hashedPwdByte := []byte(user.Password)
	pwdByte := []byte(auth.Password)
	err := bcrypt.CompareHashAndPassword(hashedPwdByte, pwdByte)
	if err != nil {
		return c.Status(403).JSON(ErrorRes{
			Message:   "Email or password is invalid",
			ErrorCode: "ERR:AUTH:INVALID",
		})
	}
	// Generate JWT
	token := GenerateJWT(user.ID)
	return c.JSON(resAuth{User: user, Token: token})
}

func Register(c *fiber.Ctx) error {
	var register userRegister
	if err := c.BodyParser(&register); err != nil {
		return c.Status(422).JSON(ErrorRes{
			Message:   "Request body INVALID",
			ErrorCode: "ERR:PARSE_JSON",
		})
	}
	// Find user with the same email in DB
	var sameEmail models.Person
	database.DB.First(&sameEmail, "email = ?", register.Email)
	if sameEmail.ID != "" {
		return c.Status(403).JSON(ErrorRes{
			Message:   "Email already registered",
			ErrorCode: "ERR:EMAIL:NOT_UNIQUE",
		})
	}
	user := models.NewUser(register.Name, register.Email, register.Password, register.Phone, register.Location)
	err := database.DB.Create(&user).Error
	if err != nil {
		log.Fatal("FAILED TO WRITE TO DB!", err.Error())
	}
	token := GenerateJWT(user.ID)
	return c.JSON(resAuth{User: user, Token: token})
}
