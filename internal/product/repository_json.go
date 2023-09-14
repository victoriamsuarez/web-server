package product

import (
	"errors"

	"github.com/victoriamsuarez/web-server/internal/domain"
	"github.com/victoriamsuarez/web-server/pkg/store"
)

// BUSCA INFORMACIÓN DE PRODUCT EN LA BASE DE DATOS

type RepositoryJson interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	GetByPrice(minPrice float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error)
	Delete(id int) error
}

type productRepositoryJson struct {
	storage store.Store
}

// NewRepository crea un nuevo repositorio
func NewRepositoryJson(storage store.Store) Repository {
	return &productRepositoryJson{storage}
}

// GetAll devuelve todos los productos
func (r *productRepositoryJson) GetAll() []domain.Product {
	products, err := r.storage.GetAll()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

// GetByID busca un producto por su id
func (r *productRepositoryJson) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.GetOne(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil
}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *productRepositoryJson) GetByPrice(price float64) []domain.Product {
	var products []domain.Product
	list, err := r.storage.GetAll()
	if err != nil {
		return products
	}
	for _, product := range list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *productRepositoryJson) Create(p domain.Product) (domain.Product, error) {
	_, err := r.existingProduct(p.Id)
	if err != nil {
		return domain.Product{}, errors.New("the product already exists")
	}

	_, err = r.validateCodeValue(p.Code_Value)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	err = r.storage.AddOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}

	return p, nil
}

func (r *productRepositoryJson) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error) {
	p := domain.Product{Name: name, Quantity: quantity, Code_Value: codeValue, Is_Published: isPublished, Expiration: expiration, Price: price}

	_, err := r.validateCodeValue(p.Code_Value)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}

	err = r.storage.UpdateOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}

func (r *productRepositoryJson) Delete(id int) error {
	err := r.storage.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepositoryJson) existingProduct(id int) (bool, error) {

	for _, uniqueId := range r.GetAll() {
		if uniqueId.Id == id {
			return false, errors.New("product already exists")
		}
	}
	return true, nil

}

func (r *productRepositoryJson) validateCodeValue(code string) (bool, error) {

	for _, uniqueCode := range r.GetAll() {
		if uniqueCode.Code_Value == code {
			return false, errors.New("product already exists")
		}
	}
	return true, nil
}
