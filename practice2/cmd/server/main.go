package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

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

// Slice de productos
var Products []Product

func main() {

	server := gin.Default()

	productPath := server.Group("/products")

	productPath.GET("/:id", GetInMemoriam)
	productPath.POST("", PostProduct)

	// Correr servidor
	server.Run(":8080")

}

func GetInMemoriam(c *gin.Context) {
	// Obtener el parámetro de la URL
	idParam := c.Param("id")

	for _, idProd := range Products {
		if strconv.Itoa(idProd.Id) == idParam {
			c.JSON(http.StatusOK, idProd)
			break
		}
	}

}

func PostProduct(c *gin.Context) {

	var productCreated Product

	if err := c.ShouldBindJSON(&productCreated); err != nil {
		c.JSON(400, "fuaaa")
		return
	}

	err := validateData(&productCreated)
	if err != nil {
		c.JSON(404, "No se pudo")
	}

	_, err = productCreated.existingProduct()
	if err != nil {
		c.JSON(404, "ya existe")
	}

	_, err = productCreated.validateCodeValue(Products)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	_, err = time.Parse("02/02/2006", productCreated.Expiration)
	if err != nil {
		c.JSON(404, "format date invalid")
		return
	}

	productCreated.Id = len(Products) + 1
	Products = append(Products, productCreated)
	c.JSON(201, productCreated)
}

func (p Product) existingProduct() (bool, error) {

	var products []Product

	for _, uniqueId := range products {
		if uniqueId.Id == p.Id {
			return false, errors.New("product already exists")
		}
	}
	return true, nil

}

func validateData(p *Product) (err error) {
	if p.Name == "" {
		err = errors.New("invalid name")
		return
	}
	if p.Identity == 0 {
		err = errors.New("invalid identity")
		return
	}
	if p.Code_value == "" {
		err = errors.New("invalid code value")
		return
	}
	if p.Expiration == "" {
		err = errors.New("invalid date")
		return
	}
	if p.Price == 0.0 {
		err = errors.New("invalid price")
		return
	}
	return
}

func (p Product) validateCodeValue(existProd []Product) (bool, error) {

	for _, uniqueCode := range existProd {
		if uniqueCode.Code_value == p.Code_value {
			return false, errors.New("product already exists")
		}
	}
	return true, nil
}
