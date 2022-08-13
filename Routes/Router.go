package Routes

import (
	"backPractica1_SO1/Controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/practice1/getCars", Controller.GetVehicles)
	app.Post("/practice1/postCars", Controller.PostVehicles)
	app.Put("/practice1/putCar/:id", Controller.PutVehicles)
	app.Delete("/practice1/delCar/:id", Controller.DelVehicles)
}
