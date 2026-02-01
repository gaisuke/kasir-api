package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (service *CategoryService) GetAll() ([]model.Category, error) {
	return service.repo.GetAll()
}

func (service *CategoryService) Create(category *model.Category) error {
	return service.repo.Create(category)
}

func (service *CategoryService) GetByID(id int) (*model.Category, error) {
	return service.repo.GetByID(id)
}

func (service *CategoryService) Update(category *model.Category) error {
	return service.repo.Update(category)
}

func (service *CategoryService) Delete(id int) error {
	return service.repo.Delete(id)
}
