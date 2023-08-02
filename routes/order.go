package routes

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

// {
// 	id: 1,
// 	user: {
// 		id: 3,
// 		first_name: "Bob",
// 		last_name: "Marley",
// 		email: "@com"
// 	},
// 	product {
// 		id: 2
// 		name: "mack"
// 		serial_number: "234"
// 		price: 200
// 	}
// }

type OrderSerializer struct {
	ID        uint               `json:"id"`
	User      UserSerializer     `json:"user"`
	Product   ProductsSerializer `json:"product"`
	CreatedAt time.Time          `json:"order_date"`
}

func CreateResponceOrder(order models.Order, user UserSerializer, product ProductsSerializer) OrderSerializer {
	return OrderSerializer{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&order)

	responceuser := CreateResponceUser(user)
	responceProduct := CreateResponceProduct(product)
	responceOrder := CreateResponceOrder(order, responceuser, responceProduct)

	return c.Status(200).JSON(responceOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Db.Find(&orders)
	responceOrders := []OrderSerializer{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.DB.Db.Find(&user, "id = ?", order.UserRefer)
		database.DB.Db.Find(&product, "id = ?", order.ProductRefer)
		responceOrder := CreateResponceOrder(order, CreateResponceUser(user), CreateResponceProduct(product))
		responceOrders = append(responceOrders, responceOrder)
	}
	return c.Status(200).JSON(responceOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.DB.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order is not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("there is no any order with this id")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(200).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.DB.Db.First(&user, order.UserRefer)
	database.DB.Db.First(&product, order.ProductRefer)
	responceUser := CreateResponceUser(user)
	responceProduct := CreateResponceProduct(product)
	responceOrder := CreateResponceOrder(order, responceUser, responceProduct)

	return c.Status(200).JSON(responceOrder)
}
