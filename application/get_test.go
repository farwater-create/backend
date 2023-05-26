// BEGIN: 8f7e6d5a9b3c
package application_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/farwater-create/backend/application"
	"github.com/farwater-create/backend/models"
	"github.com/farwater-create/backend/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGET(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := test.NewTestDB(&models.Application{})

	// Create test application
	app := &models.Application{
		Status: "pending",
		Reason: "I really want to join this project",
		UserID: 1,
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db.Create(app)

	router := test.NewTestRouter(db)

	router.GET("/applications/:id", application.GET)

	// test exist application
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/applications/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseApplication models.Application
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseApplication))
	assert.Equal(t, app.ID, responseApplication.ID)
	assert.Equal(t, app.Status, responseApplication.Status)
	assert.Equal(t, app.Reason, responseApplication.Reason)
	assert.Equal(t, app.UserID, responseApplication.UserID)
	assert.Equal(t, app.CreatedAt.Unix(), responseApplication.CreatedAt.Unix())
	assert.Equal(t, app.UpdatedAt.Unix(), responseApplication.UpdatedAt.Unix())
	assert.Equal(t, app.Status, responseApplication.Status)

	// test not exist application
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/applications/2", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
