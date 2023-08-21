package routes

import (
	"bkndOpenMind/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get(config.ConcatenateStrings(config.PathApi), FirstEP)
	app.Post(config.ConcatenateStrings(config.PathApi, config.PathU), CreateUser)

	app.Get(config.ConcatenateStrings(config.PathApi, config.PathU), GetUsers)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), GetUser)
	app.Put(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), UpdateUser)
	app.Delete(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), DeleteUser)
	// Ep for Product
	// app.Post(config.ConcatenateStrings(config.PathApi, config.PathP), CreateProduct)
	// app.Get(config.ConcatenateStrings(config.PathApi, config.PathP), GetProducts)
	// app.Get(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), GetProduct)
	// app.Put(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), UpdateProduct)
	// app.Delete(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), DeleteProduct)
	// EP for orders
	// app.Post(config.ConcatenateStrings(config.PathApi, config.PathO), CreateOrder)
	// app.Get(config.ConcatenateStrings(config.PathApi, config.PathO), GetOrders)
	// app.Get(config.ConcatenateStrings(config.PathApi, config.PathO, config.PathID), GetOrder)

}
