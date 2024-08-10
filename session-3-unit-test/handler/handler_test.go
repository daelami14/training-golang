package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"training-golang/session-3-unit-test/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHelloMessage(t *testing.T) {
	t.Run("Positive Case - Correct Message", func(t *testing.T) {
		expectedOutput := "Hallo dari Gin!"
		actualOutput := handler.GetHelloMessage()
		require.Equal(t, expectedOutput, actualOutput, "The message should be '%s'", expectedOutput)
	})
}

func TestRootHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", handler.RootHandler)

	//create a new http request
	req, _ := http.NewRequest("GET", "/", nil)

	//create a responserecorder to record the response
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	expectedBody := `{"message":"Hallo dari Gin!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
