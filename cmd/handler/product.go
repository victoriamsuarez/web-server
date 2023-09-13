package handler

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victoriamsuarez/web-server/practice4/internal/domain"
	"github.com/victoriamsuarez/web-server/practice4/internal/product"
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

func validateData(p *domain.Product) (bool, error) {
	if p.Name == "" {
		return false, errors.New("invalid name")

	}
	if p.Quantity == 0 {
		return false, errors.New("invalid identity")
	}
	if p.Code_Value == "" {
		return false, errors.New("invalid code value")
	}
	if p.Expiration == "" {
		return false, errors.New("invalid date")
	}
	if p.Price == 0.0 {
		return false, errors.New("invalid price")
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
		_, err = time.Parse("02/02/2006", product.Expiration)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "format date invalid"})
			return
		}
		valid, err := validateData(&product)
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

func (h *productHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product

		// return func(c *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validateData(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		_, err = time.Parse("02/02/2006", product.Expiration)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "format date invalid"})
			return
		}
		p, err := h.s.Update(int(id), product.Name, product.Quantity, product.Code_Value, product.Is_Published, product.Expiration, product.Price)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, p)
	}
}

func (h *productHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product

		// return func(c *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		if product.Price == 0.0 {
			ctx.JSON(400, gin.H{"error": "price is required"})
			return
		}
		p, err := h.s.UpdatePrice(int(id), product.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, p)
	}
}

func (h *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("product %d has been removed", id)})
	}
}
