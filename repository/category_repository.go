package repository

import "training-go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
