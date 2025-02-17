package tests_user

import (
	"errors"
	"testing"

	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) Login(input validation.Login) (models.User, error) {
	args := m.Called(input)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *mockUserService) Register(input validation.RegisterUserStoreAdminInput) (models.User, error) {
	args := m.Called(input)
	return args.Get(1).(models.User), args.Error(1)
}

func TestLoginUserSuccess(t *testing.T) {
	mockService := new(mockUserService)
	testInput := validation.Login{
		Email:    "khoda@mail.com",
		Password: "password",
	}

	user := models.User{
		Name:  "khoda",
		Email: "khoda@mail.com",
		Phone: "0932093232",
	}

	mockService.On("Login", testInput).Return(user, nil)

	result, err := mockService.Login(testInput)

	assert.NoError(t, err)
	assert.Equal(t, user, result)

	mockService.AssertCalled(t, "Login", testInput)
}
func TestLoginUserFailed(t *testing.T) {
	mockService := new(mockUserService)
	testInput := validation.Login{
		Email:    "khoda@mail.com",
		Password: "password",
	}

	user := models.User{
		Name:  "khoda",
		Email: "khoda@mail.com",
		Phone: "0932093232",
	}

	mockService.On("Login", testInput).Return(user, errors.New("login error"))

	result, err := mockService.Login(testInput)

	assert.Error(t, err)
	assert.Equal(t, user, result)

	mockService.AssertCalled(t, "Login", testInput)
}
