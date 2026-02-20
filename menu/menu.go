package menu

import "fmt"

func ShowMenu() {
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
			fmt.Println("test")

		}
	}
}
