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
	users       []*user.User
	currentUser *user.User
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

func (a *Auth) Login(email, password string) {
	for i := 0; i < len(a.users); i++ {
		u := a.users[i]
		if u.Email == email && u.Password == password {
			a.currentUser = u
			fmt.Println("\nLogin berhasil", u.FullName())
			return
		}
	}

	fmt.Println("\nEmail atau password salah")
}

func (a *Auth) Logout() {
	if a.currentUser == nil {
		fmt.Println("\nBelum ada user login")
		return
	}

	fmt.Println("\nLogout berhasil", a.currentUser.FullName())
	a.currentUser = nil
}

func (a *Auth) ForgotPassword(email, password string) {}
func (a *Auth) ListUsers()                            {}
func (a *Auth) IsLoggedIn() bool {
	return a.currentUser != nil
}
func (a *Auth) CurrentUser() *user.User {
	return a.currentUser
}
