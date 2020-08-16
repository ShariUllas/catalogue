package response

import "catalogue/api/internal/core/data"

// AddCategory is
type AddCategory struct {
	ID int `json:"id"`
}

// Category -
type Category struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category *string `json:"category"`
}

// NewAddCategory -
func NewAddCategory(id *int) AddCategory {
	return AddCategory{ID: *id}
}

// NewListCategoryResponse -
func NewListCategoryResponse(data []*data.Category) []*Category {
	var categories []*Category
	for _, v := range data {
		var category Category
		category.ID = v.ID
		category.Name = v.Name
		category.Category = v.Category
		categories = append(categories, &category)
	}
	return categories
}
