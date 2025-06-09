package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}
func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category

		err := rows.Scan(&category.ID, &category.Name, &category.Description)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByID(id string) (category Category, err error) {

	row := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id)

	err = row.Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return Category{}, nil // No category found
		}
		return Category{}, err // Other error
	}

	return category, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	row := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories as c JOIN courses as co ON c.id = co.category_id WHERE co.id = ?", courseID)

	var category Category

	err := row.Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return Category{}, nil // No category found
		}
		return Category{}, err // Other error
	}

	return category, nil
}
