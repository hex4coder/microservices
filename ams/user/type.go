package main

import "context"

type User struct {
	ID       string
	Email    string
	Password string
}

type UserService interface {
	CreateUser(context.Context) error
}

type UserStore interface {
	Create(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
	Get(context.Context) error
}
