package service

import "koda-b6-golang5/user"

type AuthService interface {
	Register(firstName, lastName, email, password string)
	Login(email, password string)
	Logout()
	ForgotPassword(email, password string)
	ListUsers()
	IsLoggedIn() bool
	CurrentUser() *user.User
}
