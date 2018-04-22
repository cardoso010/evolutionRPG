package utils

import (
	"../models"
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// exibir informacoes do personagem
func Perfil(person models.Person) {
	fmt.Println("-----------------------------")
	fmt.Println("Name:", person.Name)
	fmt.Println("Level:", person.Level)
	fmt.Println("Experience:", person.Exp)
	fmt.Println("Hp:", person.Hp)
	fmt.Println("Force:", person.Skill.Force)
	fmt.Println("Defense:", person.Skill.Defense)
	fmt.Println("-----------------------------")
}

// exibir informacoes da creature
func Creature(creature models.Creature){
	fmt.Println("------------------------")
	fmt.Println("Name:", creature.Name)
	fmt.Println("HP:", creature.Hp)
	fmt.Println("Force:", creature.Force)
	fmt.Println("Defense:", creature.Defense)
}
