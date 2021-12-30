package service

import (
	"testing"
	"training-go-unit-test/entity"
	"training-go-unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, error := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_GetSucces(t *testing.T) {
	// Program mock
	category := entity.Category{
		Id:   "2",
		Name: "HandPhone",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, error := categoryService.Get("2")

	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
