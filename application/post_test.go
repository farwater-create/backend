package application_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farwater-create/backend/application"
	"github.com/farwater-create/backend/models"
	"github.com/farwater-create/backend/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPOST(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := test.NewTestDB(&models.Application{})

	// Create test application
	app := &application.PostApplicationInput{
		UserID: 1,
		Reason: "I really want to join this project",
	}

	db.Create(app)

	router := test.NewTestRouter(db)

	router.POST("/applications/", application.POST)

	w := httptest.NewRecorder()

	req := test.NewJSONRequest("POST", "/applications/", gin.H{
		"userID": app.UserID,
		"reason": app.Reason,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// test if application is actually created
	dbApp := &models.Application{}
	db.First(dbApp, 1)
	assert.Equal(t, app.UserID, dbApp.UserID)
	assert.Equal(t, app.Reason, dbApp.Reason)
	assert.Equal(t, "pending", dbApp.Status)

	var responseApplication models.Application
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseApplication))
}
