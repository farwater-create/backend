// BEGIN: q6j3f9d4j3f9
package apikey

import (
	"fmt"
	"testing"

	"github.com/farwater-create/backend/apiperms"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	userID := uint(1)
	permissions := []string{apiperms.Grant, apiperms.Applications}

	apiKey, err := New(userID, permissions)

	assert.NoError(t, err)
	assert.NotNil(t, apiKey)
	assert.Equal(t, userID, apiKey.UserID)
	assert.NotEmpty(t, apiKey.Key)
	assert.Equal(t, fmt.Sprintf("%s;%s", apiperms.Grant, apiperms.Applications), apiKey.Permissions)
}
