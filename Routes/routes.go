package Routes

import (
	"User_Management/Users"

	"github.com/gofiber/fiber/v2"
)

func UsermManagement(app *fiber.App) {

	app.Post("/login", Users.Login)
	app.Post("/add/users", Users.CreateUser)
	app.Delete("/delete/user/:id", Users.DeleteUser)
	app.Get("/users/list", Users.GetListUsers)
	app.Get("/user/:id", Users.GetUser)
	app.Patch("/user/:id", Users.UpdateUser)
}
