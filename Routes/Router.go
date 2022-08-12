package Routes

import (
	"backPractica1_SO1/Controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/practice1/getCars", Controller.GetVehicles)
}
