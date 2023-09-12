package domain

// EN CASO DE DTO, USAR CARPETA pkg/util
// AGRUPA ARCHIVOS DE ESTRUCTURAS -> structs
// POR CADA DOMINIO UNA CARPETA. EN ESTE CASO product.go -> ./product -> respository.go -> service.go
// EN ESTE ARCHIVO PUEDEN ESTAR LOS GETTERS Y SETTERS

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	Code_Value   string  `json:"code_value" binding:"required"`
	Is_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}
