package main

import (
	"base_crud/app/models"
	"fmt"
)

func main() {
	u, _ := models.GetUsers()
	fmt.Println(u)
}
