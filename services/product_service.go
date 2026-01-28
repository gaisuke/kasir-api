package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (service *ProductService) GetAll() ([]model.Product, error) {
	return service.repo.GetAll()
}

func (service *ProductService) Create(product *model.Product) error {
	return service.repo.Create(product)
}

func (service *ProductService) GetByID(id int) (*model.Product, error) {
	return service.repo.GetByID(id)
}

func (service *ProductService) Update(product *model.Product) error {
	return service.repo.Update(product)
}

func (service *ProductService) Delete(id int) error {
	return service.repo.Delete(id)
}
