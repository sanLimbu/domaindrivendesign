package entity

import "github.com/google/uuid"

type Person struct {
	Name string
	ID   uuid.UUID
}
