package menu

import (
	"fmt"
	"koda-b6-golang5/service"
)

func ShowMenu(auth service.AuthService) {
	for {
		fmt.Println("\n--- Welcome to System ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Forgot Password")
		fmt.Println("0. Exit")
		fmt.Print("\nMasukkan pilihan: ")

		var code string
		fmt.Scanln(&code)

		switch code {
		case "1":
			var firstName, lastName, email, password string

			fmt.Print("First Name: ")
			fmt.Scanln(&firstName)

			fmt.Print("Last Name: ")
			fmt.Scanln(&lastName)

			fmt.Print("Email: ")
			fmt.Scanln(&email)

			fmt.Print("Password: ")
			fmt.Scanln(&password)

			auth.Register(firstName, lastName, email, password)
		}
	}
}
