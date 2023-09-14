package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victoriamsuarez/web-server/cmd/handler"
	"github.com/victoriamsuarez/web-server/internal/product"
	"github.com/victoriamsuarez/web-server/pkg/store"
)

// PUNTO DE ENTRADA
func main() {

	storage := store.NewStore("../../internal/docs/data/products.json")

	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	token := os.Getenv("TOKEN")
	repo := product.NewRepositoryJson(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service, token)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	products := r.Group("/products")

	products.GET("", productHandler.GetAllProducts())
	products.GET(":id", productHandler.GetProductById())
	products.GET("/search", productHandler.SearchProductPrice())
	products.POST("", productHandler.NewProduct())
	products.PUT(":id", productHandler.UpdateProductPut())
	products.PATCH(":id", productHandler.UpdateProductPatch())
	products.DELETE(":id", productHandler.DeleteProduct())

	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
// func loadProducts(path string, list *[]domain.Product) {
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = json.Unmarshal([]byte(file), &list)
// 	if err != nil {
// 		panic(err)
// 	}
// }
