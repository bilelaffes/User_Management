package Users

import (
	"User_Management/Auth"
	"User_Management/Database"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Permet de mettre à jour un utilisateur */
func UpdateUser(c *fiber.Ctx) error {

	err_token := Auth.CheckToken(c)
	if err_token != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Message": err_token.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	userId := c.Params("id")
	var user Database.User
	var infoUser map[string]interface{}

	defer cancel()

	/* le BodyParser et Marshal permettent d'éviter d'ajouter des champs qui n'existe pas pour un user */
	if err := c.BodyParser(&user); err != nil { // récupérer l'objet JSON du body et le stock dans la variable user
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	out, _ := json.Marshal(user) // permet de sérialiser l'objet
	if err := json.Unmarshal(out, &infoUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	for info := range infoUser {
		if infoUser[info] == "" || infoUser[info] == nil || infoUser[info] == 0.0 {
			delete(infoUser, info)
		}
	}

	options := options.Update().SetUpsert(false)
	res, err := Database.UpdateOne(ctx, bson.M{"id": userId}, bson.M{"$set": infoUser}, options)
	if res.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"Message": "User not found!"})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	data, err := ioutil.ReadFile(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
	}

	if string(data) != infoUser["data"] && infoUser["data"] != nil {
		err = os.Remove(userId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
		}

		file, err := os.OpenFile(userId, os.O_CREATE|os.O_RDWR, 0600)
		defer file.Close()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
		}

		_, err = file.WriteString(infoUser["data"].(string))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Message": err.Error()})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Message": "Updated successfully"})
}
