package controller

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/3380issei/fib_api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewFibController(t *testing.T) {
	mockFu := &usecase.MockFibUsecase{}
	fc := NewFibController(mockFu)

	assert.NotNil(t, fc)
}

func TestFibController_GetFib(t *testing.T) {
	tests := map[string]struct {
		normal         bool
		n              string
		expectedStatus int
		expectedRes    GetFibResponse
	}{
		"正常系:n=5": {
			normal:         true,
			n:              "5",
			expectedStatus: 200,
			expectedRes:    GetFibResponse{Result: "5"},
		},
		"異常系:n=0": {
			normal:         false,
			n:              "0",
			expectedStatus: 400,
			expectedRes:    GetFibResponse{Message: InvalidParameterMessageNotNatural},
		},
		"異常系:n=-1": {
			n:              "-1",
			expectedStatus: 400,
			expectedRes:    GetFibResponse{Message: InvalidParameterMessageNotNatural},
		},
		"異常系:n=abc": {
			n:              "abc",
			expectedStatus: 400,
			expectedRes:    GetFibResponse{Message: InvalidParameterMessageNotInt},
		},
	}

	for name, tt := range tests {
		mockFu := new(usecase.MockFibUsecase)

		if tt.normal {
			n, _ := strconv.Atoi(tt.n)
			mockFu.On("GetFib", uint(n)).Return(*big.NewInt(int64(n)), nil)
		}
		fc := fibController{mockFu}

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		c.Request = httptest.NewRequest("GET", fmt.Sprintf("/fib?n=%v", tt.n), nil)
		c.Request.Header.Set("Content-Type", "application/json")

		t.Run(name, func(t *testing.T) {
			fc.GetFib(c)
			jsonString := r.Body.String()

			var data GetFibResponse
			if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
				t.Errorf(err.Error())
			}

			assert := assert.New(t)
			assert.Equal(tt.expectedStatus, r.Code)
			assert.Equal(tt.expectedRes.Result, data.Result)
			assert.Equal(tt.expectedRes.Message, data.Message)
		})
	}
}
