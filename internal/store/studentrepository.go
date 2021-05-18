package store

import (
	"github.com/zn11ch/SimpleWebApp/internal/model"
)

type StudentRepository struct {
	store *Store
}

func (r *StudentRepository) Create(u *model.Student) (*model.Student, error) {
	if err := r.store.db.QueryRow("INSERT INTO students (fullname, faculty, course) VALUES ($1, $2, $3) RETURNING id",
		u.FullName,
		u.Faculty,
		u.Course,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *StudentRepository) FindById(id int) (*model.Student, error) {
	u := &model.Student{}
	if err := r.store.db.QueryRow(
		"SELECT id, fullname, faculty, course FROM students WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.FullName,
		&u.Faculty,
		&u.Course,
	); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *StudentRepository) ListAll() ([]model.Student, error) {
	var arr []model.Student
	rows, _ := r.store.db.Query(
		"SELECT id, fullname, faculty, course FROM students ORDER BY id ASC;",
	)
	for rows.Next() {
		u := &model.Student{}
		_ = rows.Scan(&u.ID,
			&u.FullName,
			&u.Faculty,
			&u.Course)
		arr = append(arr, *u)
	}
	return arr, nil
}

func (r *StudentRepository) Update(u *model.Student) (*model.Student, error) {
	if err := r.store.db.QueryRow(
		"UPDATE students SET fullname=$2, faculty=$3, course=$4 WHERE id = $1 RETURNING id",
		u.ID,
		u.FullName,
		u.Faculty,
		u.Course,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil

}
