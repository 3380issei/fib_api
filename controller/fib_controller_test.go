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
		isSuccess      bool
		n              string
		expectedStatus int
		expectedRes    interface{}
	}{
		"正常系:n=5": {
			isSuccess:      true,
			n:              "5",
			expectedStatus: 200,
			expectedRes:    GetFibResponseSuccess{Result: "5"},
		},
		"異常系:n=0": {
			isSuccess:      false,
			n:              "0",
			expectedStatus: 400,
			expectedRes:    GetFibResponseError{Message: InvalidParameterMessageNotNatural},
		},
		"異常系:n=-1": {
			n:              "-1",
			expectedStatus: 400,
			expectedRes:    GetFibResponseError{Message: InvalidParameterMessageNotNatural},
		},
		"異常系:n=abc": {
			n:              "abc",
			expectedStatus: 400,
			expectedRes:    GetFibResponseError{Message: InvalidParameterMessageNotInt},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(r)
			c.Request = httptest.NewRequest("GET", fmt.Sprintf("/fib?n=%s", tt.n), nil)
			c.Request.Header.Set("Content-Type", "application/json")

			mockFu := new(usecase.MockFibUsecase)
			if tt.isSuccess {
				n, _ := strconv.Atoi(tt.n)
				mockFu.On("GetFib", uint(n)).Return(*big.NewInt(int64(n)), nil)
			}
			fc := fibController{mockFu}
			fc.GetFib(c)

			jsonString := r.Body.String()
			var res interface{}
			if tt.isSuccess {
				var data GetFibResponseSuccess
				json.Unmarshal([]byte(jsonString), &data)
				res = data
			} else {
				var data GetFibResponseError
				json.Unmarshal([]byte(jsonString), &data)
				res = data
			}

			assert := assert.New(t)
			assert.Equal(tt.expectedStatus, r.Code)
			assert.Equal(tt.expectedRes, res)
		})
	}
}
