package apikey

import (
	"strings"

	"github.com/farwater-create/backend/models"
	"github.com/google/uuid"
)

func New(user uint, permissions []string) (*models.ApiKey, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	apiKey := &models.ApiKey{
		UserID:      user,
		Key:         u.String(),
		Permissions: strings.Join(permissions, ";"),
	}
	return apiKey, nil
}
