package main

import (
	"base_crud/api/config"
	"fmt"
	"log"
)

func main() {
	fmt.Println(config.Config.API_PORT)
	log.Print("A")
}
