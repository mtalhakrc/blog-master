package model

import "github.com/blog-master/pkg/model"

type UserType int

const (
	UserTypeNormal UserType = iota
	UserTypeAdmin
)

type User struct {
	model.BaseModel
	Name     string
	Surname  string
	Email    string
	Password string
	Type     UserType
}
