package utils

import (
	"fmt"
	"os"
	"os/exec"

	"../models"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Perfil(person models.Person) {
	fmt.Println("-----------------------------")
	fmt.Println("Name:", person.Name)
	fmt.Println("Level:", person.Level)
	fmt.Println("Experience:", person.Experience)
	fmt.Println("Hp:", person.Hp)
	fmt.Println("Force:", person.Skill.Force)
	fmt.Println("Defense:", person.Skill.Defense)
	fmt.Println("-----------------------------")
}