package models

import "github.com/google/uuid"

type Team struct {
	//A team is a kind of User with additional information about members
	User
	// Members of the team
	Members []uuid.UUID
}
