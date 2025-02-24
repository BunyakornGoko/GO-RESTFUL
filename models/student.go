package models

import "example.com/rest-api/db"

type Address struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	Zip    string `json:"zip" binding:"required"`
}

type Student struct {
	ID         int64   `json:"id"`
	Name       string `json:"name" binding:"required"`
	University string `json:"university" binding:"required"`
	Department string `json:"department" binding:"required"`
	Age        int64 `json:"age" binding:"required"`
	Address    *Address `json:"address" binding:"required"`
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

	result, err := stmt.Exec(s.Name, s.University, s.Department, s.Age, s.Address.Street, s.Address.City, s.Address.State, s.Address.Zip)
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
		err := rows.Scan(&student.ID, &student.Name, &student.University, &student.Department, &student.Age, &address.Street, &address.City, &address.State, &address.Zip)
		if err != nil {
			return nil, err
		}
		student.Address = &address
		students = append(students, student)
	}
	return students, nil
}