package models

import "github.com/google/uuid"

type Team struct {
	// The user's unique ID
	UserId uuid.UUID
	// This UploadKey allows the user to upload images under their account
	UploadKey string
	// Whether or not this user is a premium user
	Premium bool
}
