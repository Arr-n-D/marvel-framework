package model

// 2nd highest level entity inside Model

// This file can freely import anything below it

// Below it are:
// - book.go
// - page.go

type Shelf struct {
	Id    string
	Books []Book
}
