package usecase

import (
	"math/big"

	"github.com/stretchr/testify/mock"
)

type MockFibUsecase struct {
	mock.Mock
}

func (m *MockFibUsecase) GetFib(n uint) big.Int {
	arg := m.Called(n)
	return arg.Get(0).(big.Int)
}
