package model

import (
	"fmt"
)

type ErrDatabaseNotFound struct {
	entityType string
	entityId   uint
}

func (e *ErrDatabaseNotFound) Error() string {
	return fmt.Sprintf("Couldn't find %s with id %v", e.entityType, e.entityId)
}
