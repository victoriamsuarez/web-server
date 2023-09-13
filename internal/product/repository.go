package product

import (
	"errors"

	"github.com/victoriamsuarez/web-server/practice4/internal/domain"
)

// BUSCA INFORMACIÓN DE PRODUCT EN LA BASE DE DATOS

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	GetByPrice(minPrice float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error)
	UpdatePrice(id int, price float64) (domain.Product, error)
	Delete(id int) error
}

// Añadir el método PUT a nuestra API, recordemos que crea o reemplaza un recurso en su totalidad con el contenido en la request. Tené en cuenta validar los campos que se envían, como hiciste con el método POST. Seguimos aplicando los cambios sobre la lista cargada en memoria.

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

	_, err := r.existingProduct(p.Id)
	if err != nil {
		return domain.Product{}, errors.New("the product already exists")
	}

	_, err = r.validateCodeValue(p.Code_Value)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

func (r *productRepository) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error) {
	p := domain.Product{Name: name, Quantity: quantity, Code_Value: codeValue, Is_Published: isPublished, Expiration: expiration, Price: price}

	updated := false

	for i := range r.list {
		if r.list[i].Id == id {
			p.Id = id
			r.list[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, errors.New("id not found")
	}
	return p, nil
}

func (r *productRepository) UpdatePrice(id int, price float64) (domain.Product, error) {
	var p domain.Product
	updated := false
	for i := range r.list {
		if r.list[i].Id == id {
			r.list[i].Price = price
			updated = true
			p = r.list[i]
		}
	}
	if !updated {
		return domain.Product{}, errors.New("id not found")
	}
	return p, nil
}

func (r *productRepository) Delete(id int) error {
	deleted := false
	var index int
	for i := range r.list {
		if r.list[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return errors.New("id not found")
	}
	r.list = append(r.list[:index], r.list[index+1:]...)
	return nil
}

func (r *productRepository) existingProduct(id int) (bool, error) {

	for _, uniqueId := range r.list {
		if uniqueId.Id == id {
			return false, errors.New("product already exists")
		}
	}
	return true, nil

}

func (r *productRepository) validateCodeValue(code string) (bool, error) {

	for _, uniqueCode := range r.list {
		if uniqueCode.Code_Value == code {
			return false, errors.New("product already exists")
		}
	}
	return true, nil
}
