package test

import (
	"github.com/shivang-saxena/go-gen-gorm/internal/queryset/generator/tmp"
	"gorm.io/gorm"
)

//go:generate go run ../../../../cmd/go-gen-gorm/go_gen_gorm.go -in models.go

// User is a usual user
// gen:qs
type User struct {
	gorm.Model

	//Posts []Post
	Name    string
	Surname *string `gorm:"column:user_surname"`
	Email   string
}

// Blog is a blog
// gen:qs
type Blog struct {
	gorm.Model

	Name string `gorm:"column:myname"`
}

// Post is an article
// gen:qs
type Post struct {
	gorm.Model

	Blog   *Blog // may be no blog
	User   User
	Title  *string
	Str    tmp.StringDef
	Unused int `gorm:"-"`
}

// String is just for testing purposes
func (p *Post) String() string {
	return ""
}

// SomeMethod is just for testing purposes
func (b *Blog) SomeMethod() string {
	if b.ID%2 == 0 {
		return "1"
	}

	return "0"
}

// CheckReservedKeywords is a struct for checking
// work of fields with reserved keywords names
// gen:qs
type CheckReservedKeywords struct {
	Type   string
	Struct int
}
