package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type ProductsService struct {
	Repo repository.ProductsRepository
}

func NewProductsService(repo repository.ProductsRepository) *ProductsService {
	return &ProductsService{
		Repo: repo,
	}
}

func (s *ProductsService) CreateProducts(products *domain.Products) (*domain.Products, error) {
	result, err := s.Repo.CreateProducts(products)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductsService) GetProducts() (*[]domain.Products, error) {
	products, err := s.Repo.FindProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductsService) UpdateProducts(products *domain.Products) (*domain.Products, error) {
	product, err := s.Repo.UpdateProducts(products)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductsService) DeleteProducts(products *domain.Products) error {
	if err := s.Repo.DeleteProducts(products); err != nil {
		return err
	}

	return nil
}
