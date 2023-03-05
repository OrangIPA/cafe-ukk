package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	"github.com/OrangIPA/ukekehfrozekakhyr/db"
	"github.com/OrangIPA/ukekehfrozekakhyr/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(c *fiber.Ctx) error {
	// Parse body
	u := new(LoginUserParams)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Return if username or password is empty
	if u.Password == "" || u.Username == "" {
		return c.SendStatus(400)
	}

	// Get user password
	var users []model.User = nil
	if err := db.DB.Where("username = ?", u.Username).Find(&users).Error; err != nil {
		return err
	}
	if len(users) < 1 {
		return c.Status(401).Send([]byte("Username doesn't exist"))
	}
	user := users[0]
	correctPass := user.Password
	inputPass := u.Password

	// Verify the password
	h := sha256.New()
	h.Write([]byte(inputPass))
	hashedPass := h.Sum(nil)
	hexPass := hex.EncodeToString(hashedPass)
	if hexPass != correctPass {
		return c.Status(401).Send([]byte("Incorrect password"))
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"userId": user.UserID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
