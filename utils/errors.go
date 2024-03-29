package utils

import "fmt"

type ErrDatabaseNotFound struct {
	EntityType string
	EntityId   uint
}

func (e *ErrDatabaseNotFound) Error() string {
	return fmt.Sprintf("Couldn't find %s with id %v", e.EntityType, e.EntityId)
}
