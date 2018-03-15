package main

import (
	"fmt"

	"./models"
	"./utils"
)

func main() {
	char := models.Person{Level: 1, Experience: 0, Hp: 100, Skill: models.Skill{Force: 0, Defense: 0}}

	fmt.Println("Digite um nome:")
	fmt.Scan(&char.Name)

	utils.Clear()
	fmt.Println("Bem vindo ao game", char.Name, "!")
	fmt.Println("------------------")
	fmt.Println("Escolha uma mapa para jogar:")

	cave := models.Cave{}
	for i, mapa := range cave.ListCave() {
		fmt.Printf("[%d] - %s \n", i, mapa.Name)
	}

}
