package handler

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victoriamsuarez/web-server/internal/domain"
	"github.com/victoriamsuarez/web-server/internal/product"
)

// URLS DE LOS PRODUCTOS
// EN ESTE CASO NO HAY CONTROLADOR PORQUE TENGO UN HANDLER
//

type productHandler struct {
	s     product.Service
	token string
}

func NewProductHandler(s product.Service, token string) *productHandler {
	handler := &productHandler{s: s, token: token}
	return handler
}

func (h *productHandler) GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, _ := h.s.GetAll()
		ctx.JSON(200, products)
	}

}

func (h *productHandler) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetById(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(200, product)
	}
}

func (h *productHandler) SearchProductPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el parámetro de la URL
		priceQuery := ctx.Query("priceGt")

		// Convierte el string (que viene por parámetro) a un float64
		priceGt, err := strconv.ParseFloat(priceQuery, 64)
		if err != nil {
			ctx.JSON(400, "error: invalid price")
			return
		}

		products, err := h.s.GetByPrice(priceGt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "no products found"})
			return
		}

		// En caso de encontrar el producto por id devuelvo el producto guardado en la variable targetProd
		ctx.JSON(200, products)
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
func (h *productHandler) NewProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("TOKEN") != h.token {
			ctx.JSON(403, gin.H{"error": "invalid token"})
			return
		}
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

func (h *productHandler) UpdateProductPut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("TOKEN") != h.token {
			ctx.JSON(403, gin.H{"error": "invalid token"})
			return
		}
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

func (h *productHandler) UpdateProductPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("TOKEN") != h.token {
			ctx.JSON(403, gin.H{"error": "invalid token"})
			return
		}
		type Request struct {
			Name        string  `json:"name,omitempty"`
			Quantity    int     `json:"quantity,omitempty"`
			CodeValue   string  `json:"code_value,omitempty"`
			IsPublished bool    `json:"is_published,omitempty"`
			Expiration  string  `json:"expiration,omitempty"`
			Price       float64 `json:"price,omitempty"`
		}
		var req Request

		// return func(c *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}

		update := domain.Product{
			Name:         req.Name,
			Quantity:     req.Quantity,
			Code_Value:   req.CodeValue,
			Is_Published: req.IsPublished,
			Expiration:   req.Expiration,
			Price:        req.Price,
		}

		valid, err := validateData(&domain.Product{})
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Update(id, update.Name, update.Quantity, update.Code_Value, update.Is_Published, update.Expiration, update.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
		}
		ctx.JSON(200, p)
	}
}

func (h *productHandler) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("TOKEN") != h.token {
			ctx.JSON(403, gin.H{"error": "invalid token"})
			return
		}
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
