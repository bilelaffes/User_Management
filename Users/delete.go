package Users

import (
	"User_Management/Auth"
	"User_Management/Database"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

/* Permet de supprimer un utilisateur depuis son ID */
func DeleteUser(c *fiber.Ctx) error {

	err_token := Auth.CheckToken(c)
	if err_token != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": err_token.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("id")
	defer cancel()

	result, err := Database.DeleteOne(ctx, bson.M{"id": userId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"Message": "User not found!"})
	}

	err = os.Remove(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Message": "User successfully deleted"})
}
