package usecase

import (
	"web3ten0/go-gin-gorm-sample/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct{ mock.Mock }

func (m *MockUserUsecase) GetAll(name string) ([]*domain.User, error) {
	args := m.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserUsecase) GetById(id int) (*domain.User, error) {
	args := m.Called()
	return args.Get(0).(*domain.User), args.Error(1)
}
