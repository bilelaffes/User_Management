package Users

import (
	"User_Management/Auth"
	"User_Management/Database"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var userLogin Database.User
	var userInDatabase Database.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	err := Database.FindOne(ctx, bson.M{"id": userLogin.ID}).Decode(&userInDatabase)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": "Error id user"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInDatabase.Password), []byte(userLogin.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Message": "Error Password"})
	}

	token, err := Auth.GenerateToken(c, userLogin.ID, userLogin.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}
