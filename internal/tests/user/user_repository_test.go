package tests_user

import (
	"errors"
	"testing"

	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Save(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func TestFindByEmailSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	testUser := models.User{
		Id:    base.Id{ID: "f3c5792f-d368-4142-9533-674e98b685db"},
		Name:  "khoda store",
		Email: "khodastore@gmail.com",
	}

	mockRepo.On("FindByEmail", "khodastore@gmail.com").Return(testUser, nil)

	result, err := mockRepo.FindByEmail("khodastore@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, testUser, result)
	mockRepo.AssertCalled(t, "FindByEmail", "khodastore@gmail.com")
}

func TestFindByEmailFailed(t *testing.T) {
	mockRepo := new(MockUserRepository)
	testUser := models.User{
		Id:    base.Id{ID: "f3c5792f-d368-4142-9533-674e98b685db"},
		Name:  "khoda store",
		Email: "khodastore@gmail.com",
	}

	mockRepo.On("FindByEmail", "not@gmail.com").Return(testUser, errors.New("user not found"))

	result, err := mockRepo.FindByEmail("not@gmail.com")

	assert.Error(t, err)
	assert.Equal(t, testUser, result)

	mockRepo.AssertCalled(t, "FindByEmail", "not@gmail.com")
}

func TestSaveUserSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	newUser := models.User{
		Id:    base.Id{ID: "f3c5792f-d368-4142-9533-674e98b685db"},
		Name:  "khoda store",
		Email: "khodastore@gmail.com",
	}

	mockRepo.On("Save", newUser).Return(newUser, nil)

	result, err := mockRepo.Save(newUser)

	assert.NoError(t, err)
	assert.Equal(t, newUser, result)
	mockRepo.AssertCalled(t, "Save", newUser)
}
func TestSaveUserFailed(t *testing.T) {
	mockRepo := new(MockUserRepository)
	newUser := models.User{
		Id:    base.Id{ID: "f3c5792f-d368-4142-9533-674e98b685db"},
		Name:  "khoda store",
		Email: "khodastore@gmail.com",
	}

	mockRepo.On("Save", newUser).Return(newUser, errors.New("failed to insert "))

	result, err := mockRepo.Save(newUser)

	assert.Error(t, err)
	assert.Equal(t, newUser, result)
	mockRepo.AssertCalled(t, "Save", newUser)
}
