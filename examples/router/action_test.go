package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// https://circleci.com/blog/gin-gonic-testing/
func TestHandlerWithoutId(t *testing.T) {
	//expectResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	expectResponse := `{"action":"ping","message":"pong"}`
	expectStatus := http.StatusOK

	router := SetUpRouter()
	router.POST("/api:action", HandlerWithoutId)

	req, _ := http.NewRequest("POST", "/api:ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, expectStatus, w.Code)
}

func TestHandlerWithId(t *testing.T) {
	expectResponse := `{"_id":"6396d50edd380097d383aaa5","action":"ping","message":"pong"}`
	expectStatus := http.StatusOK

	router := SetUpRouter()
	router.POST("/api/:action", HandlerWithId)

	req, _ := http.NewRequest("POST", "/api/6396d50edd380097d383aaa5:ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, expectStatus, w.Code)
}

func TestHandlerWithString(t *testing.T) {
	expectResponse := `{"action":"ping","message":"pong","name":"DjanChew"}`
	expectStatus := http.StatusOK

	router := SetUpRouter()
	router.POST("/api/:action", HandlerWithString)

	req, _ := http.NewRequest("POST", "/api/DjanChew:ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, expectStatus, w.Code)
}

func TestNotfound(t *testing.T) {
	expectResponse := `{"message":"action not found"}`
	expectStatus := http.StatusNotFound

	router := SetUpRouter()
	router.POST("/api/:action", HandlerWithString)

	req, _ := http.NewRequest("POST", "/api/DjanChew:pong", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, expectStatus, w.Code)
}
