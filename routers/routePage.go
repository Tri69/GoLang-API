package routers

import(
	"github.com/gofiber/fiber/v2"
)
func routeHandlers(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx)error {
		return c.JSON(&fiber.Map{
			"statusCode" : 200,
			"Massage" : "Success Data Fetch",
			"Data" : "Hellow World",
		})
	})
}