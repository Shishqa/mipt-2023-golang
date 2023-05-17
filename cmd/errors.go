package main

import "fmt"

type NotFoundError struct {
	id int
}

type ConflictError struct {
	id int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Pokemon with id %d not found", e.id)
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("Pokemon with id %d already exists", e.id)
}
