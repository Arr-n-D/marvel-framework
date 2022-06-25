package model

// Highest level entity inside Model

// This file can freely import anything below it

// Below it are:
// - shelf.go
// - book.go
// - page.go

type Library struct {
	Id      string
	Name    string
	Shelves []Shelf
}
