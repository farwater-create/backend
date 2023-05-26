package application

import "github.com/farwater-create/backend/models"

type PostApplicationInput struct {
	UserID uint   `json:"userId" validate:"required"`
	Reason string `json:"reason" validate:"required"`
}

func New(applicationInput *PostApplicationInput) *models.Application {
	return &models.Application{
		UserID: applicationInput.UserID,
		Status: "pending",
		Reason: applicationInput.Reason,
	}
}
