package database

import (
	"program/model"
)

type StudentService interface {
	SaveStudent(std *model.Student) error
	UpdateStudent(id int, std *model.Student) error
	DeleteStudent(id int) error
	GetStudent(id int) (*model.Student, error)
	ListStudents() ([]*model.Student, error)
}
