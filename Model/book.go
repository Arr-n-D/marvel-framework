package model

// 3rd highest level entity inside Model

// This file can freely import anything below it

// Below it there is:
// - page.go

type Book struct {
	Id     string
	Title  string
	Author string
	Pages  []Page
}
