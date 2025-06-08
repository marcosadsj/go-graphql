package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) Create(title, description, categoryID string) (Course, error) {

	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, title, description, category_id) VALUES(?, ?, ?, ?)", id, title, description, categoryID)

	if err != nil {
		return Course{}, err
	}

	return Course{
		ID:          id,
		Title:       title,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, title, description, category_id FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var course Course

		err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
