package main

import (
	"koda-b6-golang5/menu"
	"koda-b6-golang5/service"
)

func main() {
	auth := service.NewAuth()
	menu.ShowMenu(auth)
}
