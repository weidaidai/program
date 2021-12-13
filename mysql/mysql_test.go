package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)
// a successful case

	func TestUpdaterow(t *testing.T) {
		db, mock, _ := sqlmock.New()
		mock.ExpectExec("update set person").WillReturnResult(sqlmock.NewResult(100, 1))

		id, err := Updaterow(db, 1)
		assert.Equal(t, id, int64(100))
		assert.Nil(t, err)
	}