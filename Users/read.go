package Users

import (
	"User_Management/Auth"
	"User_Management/Database"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

/* Permet de récupérer la liste des utilisateurs */
func GetListUsers(c *fiber.Ctx) error {

	err_token := Auth.CheckToken(c)
	if err_token != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": err_token.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []Database.User
	defer cancel()

	results, err := Database.GetUsersCollection("users").Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var user Database.User
		if err = results.Decode(&user); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
		}
		users = append(users, user)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Results": users})
}

/* Permet de récupérer un utilisateur */
func GetUser(c *fiber.Ctx) error {

	err_token := Auth.CheckToken(c)
	if err_token != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": err_token.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("id")
	var user Database.User
	defer cancel()

	err := Database.GetUsersCollection("users").FindOne(ctx, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"User": user})
}
