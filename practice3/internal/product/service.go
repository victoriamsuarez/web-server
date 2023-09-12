package product

import (
	"errors"

	"github.com/victoriamsuarez/web-server/practice3/internal/domain"
)

// RECIBE INFORMACION DEL HANDLER Y LA RESUELVE

// Inyección de dependencias
type productService struct {
	repository Repository
}

type Service interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	GetByPrice(minPrice float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
}

func NewService(repo Repository) Service {
	return &productService{repo}
}

// GetProducts muestra los datos de todos los productos
func (s *productService) GetAll() ([]domain.Product, error) {
	list := s.repository.GetAll()
	return list, nil
}

// GetProductsById muestra un producto que buscamos mediante el parámetro de la URL
func (s *productService) GetById(id int) (domain.Product, error) {

	p, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *productService) GetByPrice(price float64) ([]domain.Product, error) {
	list := s.repository.GetByPrice(price)
	if len(list) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return list, nil
}

func (s *productService) Create(p domain.Product) (domain.Product, error) {
	p, err := s.repository.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}
