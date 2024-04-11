package usecase

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFibUsecase(t *testing.T) {
	fu := NewFibUsecase()
	assert.NotNil(t, fu)
}

func TestFibUsecase_GetFib(t *testing.T) {
	fibNum99th, _ := new(big.Int).SetString("218922995834555169026", 10)
	tests := map[string]struct {
		n           uint
		expectedFib big.Int
	}{
		"n=1": {
			n:           1,
			expectedFib: *big.NewInt(1),
		},
		"n=2": {
			n:           2,
			expectedFib: *big.NewInt(1),
		},
		"n=10": {
			n:           10,
			expectedFib: *big.NewInt(55),
		},
		"n=99": {
			n:           99,
			expectedFib: *fibNum99th,
		},
	}

	fu := fibUsecase{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := fu.GetFib(tt.n)
			assert.Equal(t, tt.expectedFib, actual)
		})
	}
}
