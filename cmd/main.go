package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Identity     int     `json:"identity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_publish"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func main() {

	server := gin.Default()

	server.GET("/ping", GetPong)
	server.GET("/products", GetProducts)
	server.GET("/products/:id", GetProductsById)
	server.GET("/products/search", GetProductsPrice)

	// Correr servidor
	server.Run(":8080")

}

func GetPong(c *gin.Context) {
	c.JSON(200, "pong")
}

// GetProducts muestra los datos de todos los productos
func GetProducts(c *gin.Context) {

	// Abre el archivo
	file, err := os.Open("./products.json")
	if err != nil {
		c.JSON(500, "internal server error")
	}
	defer file.Close()

	// Slice de productos
	var products []Product

	// Decodifica el archivo JSON y lo convierte a la estructura de datos para manipular los datos
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		c.JSON(500, "internal server error")
	}

	// Devuelve todos los productos del slice
	c.JSON(200, products)

}

// GetProductsById muestra un producto que buscamos mediante el parámetro de la URL
func GetProductsById(c *gin.Context) {

	// Obtener el parámetro de la URL
	idParam := c.Param("id")

	// Convierte el string (que viene por parámetro) a un int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "error: invalid id")
		return
	}

	// Abre el archivo
	file, err := os.Open("./products.json")
	if err != nil {
		c.JSON(500, "internal server error")
	}
	defer file.Close()

	// Slice de productos
	var products []Product

	// Decodifica el archivo JSON y lo convierte a la estructura de datos para manipular los datos
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		c.JSON(500, "internal server error")
	}

	// Recorre el slice de productos, compara el id de cada producto por el id que le llega por parámetro y se guarda en una nueva variable
	var targetProd Product
	for _, idProd := range products {
		if idProd.Id == id {
			targetProd = idProd
			break
		}
	}

	// En caso que no se encuentre un producto con su id devuelve mensaje de error
	if targetProd.Id == 0 {
		c.JSON(404, "error: product not found")
		return
	}

	// En caso de encontrar el producto por id devuelvo el producto guardado en la variable targetProd
	c.JSON(200, targetProd)
}

// Crear una ruta /products/search que nos permita buscar por parámetro los
// productos cuyo precio sean mayor a un valor priceGt.
func GetProductsPrice(c *gin.Context) {

	// Obtener el parámetro de la URL
	priceQuery := c.Query("priceGt")

	// Convierte el string (que viene por parámetro) a un int
	priceGt, err := strconv.ParseFloat(priceQuery, 64)
	if err != nil {
		c.JSON(400, "error: invalid price")
		return
	}

	// Abre el archivo
	file, err := os.Open("./products.json")
	if err != nil {
		c.JSON(500, "internal server error")
	}
	defer file.Close()

	// Slice de productos
	var products []Product

	// Decodifica el archivo JSON y lo convierte a la estructura de datos para manipular los datos
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		c.JSON(500, "internal server error")
	}

	// Recorre el slice de productos, compara el id de cada producto por el id que le llega por parámetro y se guarda en una nueva variable
	var productsPrice []Product
	for _, priceProd := range products {
		if priceProd.Price > priceGt {
			productsPrice = append(productsPrice, priceProd)
		}
	}

	// En caso de encontrar el producto por id devuelvo el producto guardado en la variable targetProd
	c.JSON(200, productsPrice)

}
