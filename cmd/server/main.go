package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/victoriamsuarez/web-server/practice4/cmd/handler"
	"github.com/victoriamsuarez/web-server/practice4/internal/domain"
	"github.com/victoriamsuarez/web-server/practice4/internal/product"
)

// PUNTO DE ENTRADA
func main() {
	var productsList = []domain.Product{}
	loadProducts("../../internal/docs/data/products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GeAll())
		products.GET(":id", productHandler.GetById())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.PUT(":id", productHandler.Put())
		products.PATCH(":id", productHandler.Patch())
		products.DELETE(":id", productHandler.Delete())
	}
	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}