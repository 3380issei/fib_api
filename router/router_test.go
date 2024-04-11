package router

import (
	"strings"
	"testing"

	"github.com/3380issei/fib_api/controller"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	mockFc := controller.MockFibController{}

	routingInfo := []struct {
		method  string
		path    string
		handler string
	}{
		{"GET", "/fib", "GetFib"},
	}

	r := NewRouter(&mockFc)
	routes := r.Routes()

	for i, route := range routes {
		assert := assert.New(t)
		assert.Equal(routingInfo[i].method, route.Method)
		assert.Equal(routingInfo[i].path, route.Path)

		handlerName := route.Handler
		parts := strings.Split(handlerName, "controller.FibController.")
		if len(parts) < 2 {
			t.Fatalf("unexpected handler name: %s", handlerName)
		}
		handlerName = parts[1]
		handlerName = strings.TrimSuffix(handlerName, "-fm")
		assert.Equal(routingInfo[i].handler, handlerName)
	}
}
