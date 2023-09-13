package product

import (
	"errors"

	"github.com/victoriamsuarez/web-server/practice4/internal/domain"
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
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error)
	UpdatePrice(id int, price float64) (domain.Product, error)
	Delete(id int) error
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

func (s *productService) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error) {
	var p domain.Product
	p, err := s.repository.Update(id, name, quantity, codeValue, isPublished, expiration, price)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *productService) UpdatePrice(id int, price float64) (domain.Product, error) {
	var p domain.Product
	p, err := s.repository.UpdatePrice(id, price)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *productService) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
