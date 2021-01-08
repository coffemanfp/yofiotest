package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/coffemanfp/yofiotest/api/handlers"
	"github.com/coffemanfp/yofiotest/assigners"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAssignment(t *testing.T) {
	var fakeDB fakeAssignersDB
	var investment int32 = 3000
	var expectedAssigner assigners.Assigner

	expectedAssigner.Assign(investment)

	r := gin.Default()
	r.POST("/", handlers.CreateAssignment(fakeDB))

	rBody, _ := json.Marshal(gin.H{
		"investment": investment,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewReader(rBody))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var newA assigners.Assigner
	err := json.Unmarshal(w.Body.Bytes(), &newA)
	assert.Nil(t, err)

	assert.Equal(t, investment, newA.Investment)
	assert.Equal(t, expectedAssigner.Success, newA.Success)
	assert.Equal(t, expectedAssigner.CreditType300, newA.CreditType300)
	assert.Equal(t, expectedAssigner.CreditType500, newA.CreditType500)
	assert.Equal(t, expectedAssigner.CreditType700, newA.CreditType700)
}

func TestFailCreateAssignment(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		code         int
		errorMessage string
		fakeDB       fakeAssignersDB
	}{
		{
			name:         "without content type header",
			code:         http.StatusBadRequest,
			errorMessage: "invalid body",
		},
		{
			name: "invalid investment",
			body: `
				{
					"investment": 0
				}
			`,
			code:         http.StatusBadRequest,
			errorMessage: "invalid body",
		},
		{
			name: "internal server error",
			body: `
				{
					"investment": 100
				}
			`,
			code:         http.StatusInternalServerError,
			errorMessage: "Oops!",
			fakeDB: fakeAssignersDB{
				err: errors.New("fake error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.POST("/", handlers.CreateAssignment(tt.fakeDB))

			w := httptest.NewRecorder()

			var req *http.Request
			if tt.body != "" {
				req, _ = http.NewRequest("POST", "/", strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, _ = http.NewRequest("POST", "/", nil)
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.code, w.Code)

			resError := make(map[string]string)
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &resError))
			assert.Equal(t, tt.errorMessage, resError["error"])
		})
	}
}

func TestGetStats(t *testing.T) {
	var fakeDB fakeAssignersDB

	r := gin.Default()
	r.POST("/", handlers.GetStats(fakeDB))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var stats assigners.Stats
	assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &stats))
}

func TestFailGetStats(t *testing.T) {
	fakeDB := fakeAssignersDB{
		err: errors.New("fake err"),
	}

	r := gin.Default()
	r.POST("/", handlers.GetStats(fakeDB))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	resError := make(map[string]string)
	assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &resError))
	assert.Equal(t, "Oops!", resError["error"])
}

type fakeAssignersDB struct {
	err error
}

func (f fakeAssignersDB) Create(a assigners.Assigner) (newA assigners.Assigner, err error) {
	if f.err != nil {
		err = f.err
		return
	}
	return
}

func (f fakeAssignersDB) GetStats() (s assigners.Stats, err error) {
	if f.err != nil {
		err = f.err
		return
	}
	return
}
