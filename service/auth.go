package service

import (
	"fmt"
	"koda-b6-golang5/user"
)

type AuthService interface {
	Register(firstName, lastName, email, password string)
	Login(email, password string)
	Logout()
	ForgotPassword(email, password string)
	ListUsers()
	IsLoggedIn() bool
	CurrentUser() *user.User
}

type Auth struct {
	users []*user.User
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Register(firstName, lastName, email, password string) {
	newUser := &user.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	a.users = append(a.users, newUser)

	fmt.Println("\nRegistrasi berhasil! Selamat datang,", newUser.FullName())
}

func (a *Auth) Login(email, password string)          {}
func (a *Auth) Logout()                               {}
func (a *Auth) ForgotPassword(email, password string) {}
func (a *Auth) ListUsers()                            {}
func (a *Auth) IsLoggedIn() bool                      { return false }
func (a *Auth) CurrentUser() *user.User               { return nil }
