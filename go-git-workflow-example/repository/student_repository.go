package repository

import (
	"database/sql"
	"errors"

	"workflow-example.com/model"
)

type Repository struct {
	DB *sql.DB
}

// StudentRepository Repository Interface for student repository
type StudentRepository interface {
	GetAll() ([]model.Student, error)
	GetByID(id string) (*model.Student, error)
	Create(student model.Student) error
	Update(student model.Student) error
	Delete(id string) error
}

// GetAll Implementation of Repository Interface
func (r *Repository) GetAll() ([]model.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	// fallback to mock data if DB is empty
	if len(students) == 0 {
		students = []model.Student{
			{Id: "S001", Name: "Emmy", Major: "Computer Science", GPA: 3.00},
			{Id: "S002", Name: "Tammy", Major: "Computer Science", GPA: 3.50},
		}
	}

	return students, nil
}

// GetByID fetch a student by ID
func (r *Repository) GetByID(id string) (*model.Student, error) {
	var s model.Student
	err := r.DB.QueryRow("SELECT id, name, major, gpa FROM students WHERE id = ?", id).
		Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // not found
		}
		return nil, err
	}
	return &s, nil
}

// Create insert a new student
func (r *Repository) Create(student model.Student) error {
	_, err := r.DB.Exec(
		"INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
		student.Id, student.Name, student.Major, student.GPA,
	)
	return err
}

// Update an existing student
func (r *Repository) Update(student model.Student) error {
	res, err := r.DB.Exec(
		"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
		student.Name, student.Major, student.GPA, student.Id,
	)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("no student found to update")
	}
	return nil
}

// Delete a student by ID
func (r *Repository) Delete(id string) error {
	res, err := r.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("no student found to delete")
	}
	return nil
}
