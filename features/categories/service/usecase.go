package service

import (
	"bayareen-backend/features/categories"

	"github.com/go-playground/validator/v10"
)

type categoryUsecase struct {
	categoryData categories.Data
	validator    *validator.Validate
}

func NewCategoryUsecase(categoryData categories.Data) categories.Business {
	return &categoryUsecase{
		categoryData: categoryData,
		validator:    validator.New(),
	}
}

func (cu *categoryUsecase) Create(core categories.Core) (resp categories.Core, err error) {
	if err = cu.validator.Struct(core); err != nil {
		return categories.Core{}, err
	}

	resp, err = cu.categoryData.Create(core)

	if err != nil {
		return categories.Core{}, err
	}

	return resp, nil
}

func (cu *categoryUsecase) GetAll() []categories.Core {
	return cu.categoryData.GetAll()
}