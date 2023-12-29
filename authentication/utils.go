package authentication

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	user.Password = string(hashedPassword)
	if err := db.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).SendString("Email already exists")
	}

	db.Create(user)
	return c.JSON(user)
}

func Login(db *gorm.DB, c *fiber.Ctx) error {
	input := new(User)
	user := new(User)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Where("email = ?", input.Email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Incorrect password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["role"] = "user"

	if err := godotenv.Load(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "success"})
}

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Extract the role from the claims
	role, ok := claims["role"].(string)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if role != "user" {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Next()
}
