package apikey

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGET(t *testing.T) {
	// Set up Gin router and database
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&models.ApiKey{})

	// Set up test data
	userID := uint(1)
	permissions := []string{"read", "write"}
	getApiKeyInput := &GetApiKeyInput{
		UserID:      userID,
		Permissions: permissions,
	}
	jsonData, err := json.Marshal(getApiKeyInput)
	assert.NoError(t, err)

	// Create request
	req, err := http.NewRequest("GET", "/apikey", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder and handle request
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Set("db", db)
	ctx.Request = req
	GET(ctx)

	// Check response
	assert.Equal(t, http.StatusOK, rr.Code)
	var apiKey models.ApiKey
	err = json.Unmarshal(rr.Body.Bytes(), &apiKey)
	assert.NoError(t, err)
	assert.Equal(t, userID, apiKey.UserID)
	assert.Equal(t, strings.Join(permissions, ";"), apiKey.Permissions)
}
