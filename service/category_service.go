package service

import (
	"errors"
	"training-go-unit-test/entity"
	"training-go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("CATEGORY NOT FOUND")
	} else {
		return category, nil
	}
}
