package routes

// import (
// 	"bkndOpenMind/database"
// 	"bkndOpenMind/models"
// 	"errors"

// 	"github.com/gofiber/fiber/v2"
// )

// type ProductsSerializer struct {
// 	ID           uint   `json:"id" gorm:"primaryKey"`
// 	Name         string `json:"name"`
// 	SerialNumber string `json:"serial_number"`
// 	Price        int16  `json:"price"`
// }

// func CreateResponceProduct(productModel models.Product) ProductsSerializer {
// 	return ProductsSerializer{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber, Price: productModel.Price}
// }

// // create product
// func CreateProduct(c *fiber.Ctx) error {
// 	var product models.Product

// 	if err := c.BodyParser(&product); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	database.DB.Db.Create(&product)

// 	responceProduct := CreateResponceProduct(product)
// 	return c.Status(200).JSON(responceProduct)
// }

// // get products
// func findProduct(id int, product *models.Product) error {
// 	database.DB.Db.Find(&product, "id=?", id)
// 	if product.ID == 0 {
// 		return errors.New("product is not exist")
// 	}
// 	return nil
// }

// func GetProducts(c *fiber.Ctx) error {
// 	products := []models.Product{}
// 	database.DB.Db.Find(&products)
// 	responceProducts := []ProductsSerializer{}
// 	for _, product := range products {
// 		responceProduct := CreateResponceProduct(product)
// 		responceProducts = append(responceProducts, responceProduct)
// 	}

// 	return c.Status(200).JSON(responceProducts)
// }

// // get one product
// func GetProduct(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var product models.Product
// 	if err != nil {
// 		return c.Status(400).JSON("There is no any product with this id")
// 	}
// 	if err := findProduct(id, &product); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	responceProduct := CreateResponceProduct(product)
// 	return c.Status(200).JSON(responceProduct)
// }

// // update Product
// func UpdateProduct(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var product models.Product
// 	if err != nil {
// 		return c.Status(400).JSON("there is no any product with this id")
// 	}
// 	if err := findProduct(id, &product); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	type UpdateProduct struct {
// 		Name         string `json:"name"`
// 		SerialNumber string `json:"serial_number"`
// 		Price        int16  `json:"price"`
// 	}
// 	var updateData UpdateProduct

// 	if err := c.BodyParser(&updateData); err != nil {
// 		return c.Status(500).JSON(err.Error())
// 	}

// 	product.Name = updateData.Name
// 	product.SerialNumber = updateData.SerialNumber
// 	product.Price = updateData.Price
// 	database.DB.Db.Save(&product)

// 	responceProduct := CreateResponceProduct(product)
// 	return c.Status(200).JSON(responceProduct)
// }

// // delete product
// func DeleteProduct(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")

// 	var product models.Product
// 	if err != nil {
// 		return c.Status(400).JSON("There is no any products with this id")
// 	}
// 	if err := findProduct(id, &product); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	if err := database.DB.Db.Delete(&product).Error; err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	return c.Status(200).SendString("Delete is succefully")
// }
