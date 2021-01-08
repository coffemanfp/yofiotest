package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/yofiotest/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNoRoute(t *testing.T) {
	r := gin.Default()
	r.NoRoute(handlers.NoRoute)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	resError := make(map[string]string)
	assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &resError))
	assert.Equal(t, "not found", resError["error"])
}
