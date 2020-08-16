package data

import (
	"catalogue/api/internal/core/request"
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// CategoryRepo is repo interface for category
type CategoryRepo interface {
	GetCategory(ctx context.Context) ([]*Category, error)
	AddCategory(ctx context.Context, variant *request.Category) (*int, error)
}

type categoryImpl struct {
	*sql.DB
}

// NewCategoryRepo is
func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return &categoryImpl{DB: db}
}

// Category -
type Category struct {
	ID       int
	Name     string
	Category *string
}

const addCategoryQuery = `
INSERT INTO category
(
	name,
	parent_id,
	created_at,
	updated_at
)
VALUES
(
	$1,$2,$3,$4
)
returning id;
`

func (i *categoryImpl) AddCategory(ctx context.Context, category *request.Category) (*int, error) {
	var categoryID *int
	err := i.QueryRowContext(ctx, addCategoryQuery, category.Name, category.ParentID, time.Now(), time.Now()).Scan(&categoryID)
	if err != nil {
		return nil, errors.Wrapf(err, "adding category to db failed")
	}
	return categoryID, nil
}

const getCategoryQuery = `
SELECT c.id,
       c.name,
       p.name AS category
FROM   category AS c
       LEFT JOIN category AS p
              ON c.parent_id = p.id  
`

func (i *categoryImpl) GetCategory(ctx context.Context) ([]*Category, error) {
	var categories []*Category
	rows, err := i.QueryContext(ctx, getCategoryQuery)
	if err != nil {
		return nil, errors.Wrapf(err, "get category query failed")
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Category)
		if err != nil {
			return nil, errors.Wrapf(err, "get category query scan failed")
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
