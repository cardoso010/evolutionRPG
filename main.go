package main

import (
	"./models"
	"./utils"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	char := models.Person{Level: 1, Exp: 0, Hp: 100, Skill: models.Skill{Force: 0, Defense: 0}}

	fmt.Println("Digite um nome:")
	fmt.Scan(&char.Name)

	utils.Clear()
	fmt.Println("Bem vindo ao game", char.Name, "!")
	fmt.Println("------------------")
	fmt.Println("Escolha uma mapa para jogar:")

	caveModel := models.Cave{}
	caves := caveModel.ListCave()
	for i, mapa := range caves {
		fmt.Printf("[%d] - %s \n", i, mapa.Name)
	}

	var caveSelect int
	fmt.Scan(&caveSelect)
	mapaEscolhido := caves[caveSelect]

	fmt.Println(mapaEscolhido.Name)

	fmt.Println("Creature:")
	for _, creature := range mapaEscolhido.Creature {
		fmt.Println("------------------------")
		fmt.Println("Name:", creature.Name)
		fmt.Println("HP:", creature.Hp)
		fmt.Println("Force:", creature.Force)
		fmt.Println("Defense:", creature.Defense)
	}
	fmt.Println("------------------------")

	start := time.Now()
	stop := start.Add(time.Minute * 10)
	for {
		if time.Now().Minute() >= stop.Minute() || char.Hp <= 0 {
			break
		}
		monster := mapaEscolhido.Creature[rand.Intn(len(mapaEscolhido.Creature)-1)]
		models.Fight(&char, &monster)
	}

	fmt.Println("Exp:", char.Exp)

}
