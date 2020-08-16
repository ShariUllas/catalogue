package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// ErrInvalidCategory return when category name is invalid
	ErrInvalidCategory = fmt.Errorf("invalid category name")
)

// Category - requet model for variant
type Category struct {
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}

// NewAddCategory - parses request for add variants
func NewAddCategory(r *http.Request) (*Category, error) {
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		return nil, ErrInvalidJSON
	}
	err = category.validateCategory()
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (category *Category) validateCategory() error {
	if len(category.Name) == 0 {
		return ErrInvalidCategory
	}
	return nil
}
