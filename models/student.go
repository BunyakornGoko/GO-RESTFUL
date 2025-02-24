package models

import (
	"database/sql"

	"example.com/rest-api/db"
)

type Address struct {
	Street *string `json:"street" binding:"required"`
	City   *string `json:"city" binding:"required"`
	State  *string `json:"state" binding:"required"`
	Zip    *string `json:"zip" binding:"required"`
}

type Student struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	University string  `json:"university"`
	Department string  `json:"department"`
	Age        int     `json:"age"`
	Address    *Address `json:"address"`
}

func (s *Student) Save() error {
	query := `
	INSERT INTO students 
	(name, university, department, age, street, city, state, zip)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var result sql.Result

	if s.Address == nil {
		result, err = stmt.Exec(s.Name, s.University, s.Department, s.Age, nil, nil, nil, nil)
	} else {
		result, err = stmt.Exec(s.Name, s.University, s.Department, s.Age, s.Address.Street, s.Address.City, s.Address.State, s.Address.Zip)
	}

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	s.ID = id
	return nil
}

func GetAllStudents() ([]Student, error) {
	query := `SELECT * FROM students`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		var address Address
		var street, city, state, zip sql.NullString // Use sql.NullString for nullable fields
		err := rows.Scan(&student.ID, &student.Name, &student.University, &student.Department, &student.Age, &street, &city, &state, &zip)
		if err != nil {
			return nil, err
		}

		// Assign values to address fields
		if street.Valid {
			address.Street = &street.String
		}
		if city.Valid {
			address.City = &city.String
		}
		if state.Valid {
			address.State = &state.String
		}
		if zip.Valid {
			address.Zip = &zip.String
		}
		student.Address = &address
		students = append(students, student)
	}
	
	return students, nil
}