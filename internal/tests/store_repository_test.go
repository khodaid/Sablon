package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewStoreRepository(gormDB)

	store := models.Store{
		Name:       "Test Store",
		Address:    "123 Test Address",
		Phone:      "3243242424",
		Email:      "tes@mail.com",
		SupplierId: "abc-123",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"store\"").WithArgs(
		sqlmock.AnyArg(),
		store.Name,
		store.Address,
		store.Phone,
		store.Email,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		store.SupplierId,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := repo.Save(store)

	assert.NoError(t, err)
	assert.Equal(t, store.Name, result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
