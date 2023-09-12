package product

import (
	"errors"

	"github.com/victoriamsuarez/web-server/practice3/internal/domain"
)

// BUSCA INFORMACIÓN DE PRODUCT EN LA BASE DE DATOS

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	GetByPrice(minPrice float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
}

type productRepository struct {
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &productRepository{list}
}

// GetAll devuelve todos los productos
func (r *productRepository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *productRepository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *productRepository) GetByPrice(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *productRepository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.Code_Value) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *productRepository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.Code_Value == codeValue {
			return false
		}
	}
	return true
}
