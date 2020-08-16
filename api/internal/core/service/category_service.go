package service

import (
	"catalogue/api/internal/core/data"
	"catalogue/api/internal/core/request"
	"catalogue/api/internal/core/response"
	"context"
)

// Category is service interface
type Category interface {
	GetCategory(ctx context.Context) ([]*response.Category, error)
	AddCategory(ctx context.Context, category *request.Category) (response.AddCategory, error)
}

type categoryImpl struct {
	data.CategoryRepo
}

// NewCategoryService is
func NewCategoryService(categoryRepo data.CategoryRepo) Category {
	return &categoryImpl{CategoryRepo: categoryRepo}
}

func (i *categoryImpl) AddCategory(ctx context.Context, category *request.Category) (response.AddCategory, error) {
	categoryID, err := i.CategoryRepo.AddCategory(ctx, category)
	if err != nil {
		return response.AddCategory{}, err
	}
	return response.NewAddCategory(categoryID), nil
}

func (i *categoryImpl) GetCategory(ctx context.Context) ([]*response.Category, error) {
	categories, err := i.CategoryRepo.GetCategory(ctx)
	if err != nil {
		return nil, err
	}
	return response.NewListCategoryResponse(categories), nil
}
