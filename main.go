package main

import (
	"base_crud/app/models"
	"fmt"
)

func main() {
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
