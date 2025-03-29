package main
import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	//"Golang/routers"
)

type Person struct{
	id int
	name string
	class string
	nis int
}

func main() {
	appRoute := fiber.New()
	dataPerson := [...]Person{
		{
			id : 1,
			name : "Ilham Trihandoyo",
			class : "XI TKRO B",
			nis : 9081,
		},
		{
			id : 2,
			name : "Dinda Saputri",
			class : "XI TJKT A",
			nis : 9252,
		},
		{
			id : 3,
			name : "Pahsa Mulyanto Utomo",
			class : "X TBSM C",
			nis : 9774,
		},
	}
	appRoute.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello Test Api")
	})
	appRoute.Get("/api/v2/", func(c *fiber.Ctx)error {
	
		return c.JSON(&fiber.Map{
			"Data" : dataPerson,
			"Massage" : "Success Data Fetch",
			"statusCode" : 200,
		})
	})
	//routers.routeHandlers(appRoute)
	fmt.Println("yes")
	appRoute.Listen(":3006")
}