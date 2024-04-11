package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockFibController struct {
	mock.Mock
}

func (fc *MockFibController) GetFib(c *gin.Context) {
	fc.Called(c)
}
