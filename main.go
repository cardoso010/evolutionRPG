package main

import (
	"fmt"

	"./models"
)

func main() {
	cave := models.Cave{}
	cave.ListCave()
	fmt.Println(cave.ListCave())
}
