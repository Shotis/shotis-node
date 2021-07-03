// This file holds the model objects for User models
package models

import (
	"github.com/google/uuid"
)

type User struct {
	// The user's unique ID
	UserId uuid.UUID
	// This UploadKey allows the user to upload images under their account
	UploadKey string
	// Whether or not this user is a premium user
	Premium bool
}
