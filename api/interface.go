package api

import (
	"program/model"
)

type StudentService interface {
	SaveStudent(std *model.Student) error
	UpdateStudent(std *model.Student) error
	DeleteStudent(id int) error
	GetStudent(id int) (*model.Student, error)
	ListStudents() ([]*model.Student, error)
}
