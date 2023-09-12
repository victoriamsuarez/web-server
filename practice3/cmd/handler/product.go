package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/victoriamsuarez/web-server/practice3/internal/domain"
	"github.com/victoriamsuarez/web-server/practice3/internal/product"
)

// URLS DE LOS PRODUCTOS
// EN ESTE CASO NO HAY CONTROLADOR PORQUE TENGO UN HANDLER
//

type productHandler struct {
	s product.Service
}

func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

func (h *productHandler) GeAll() gin.HandlerFunc {

	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		c.JSON(200, products)
	}

}

func (h *productHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetById(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		c.JSON(200, product)
	}
}

func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el parámetro de la URL
		priceQuery := c.Query("priceGt")

		// Convierte el string (que viene por parámetro) a un float64
		priceGt, err := strconv.ParseFloat(priceQuery, 64)
		if err != nil {
			c.JSON(400, "error: invalid price")
			return
		}

		products, err := h.s.GetByPrice(priceGt)
		if err != nil {
			c.JSON(404, gin.H{"error": "no products found"})
			return
		}

		// En caso de encontrar el producto por id devuelvo el producto guardado en la variable targetProd
		c.JSON(200, products)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.Code_Value == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(product *domain.Product) (bool, error) {
	dates := strings.Split(product.Expiration, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

// Post crear un producto nuevo
func (h *productHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = validateExpiration(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}
