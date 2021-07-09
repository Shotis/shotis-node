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
	//StorageUsed is the amount of storage currently being used by the user in bytes
	StorageUsed float64
	//StorageLimit is the amount storage available to the user
	StorageLimit float64
	// Whether or not this user is a premium user
	Premium bool
}
