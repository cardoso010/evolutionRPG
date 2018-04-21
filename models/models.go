package models

import (
	"fmt"
	"math/rand"
)

// Action responsavel pelos metodos de acao
type Action interface {
	Attack(attacked *struct{})
	Defend() bool
}

// Skill  struct responsavel skills
type Skill struct {
	Force   int `json:"force"`
	Defense int `json:"defense"`
}

// Person  struct responsavel pelos personagem
type Person struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Level    int     `json:"level"`
	Exp      float64 `json:"experience"`
	Skill
	Hp int `json:"hp"`
}

// Attack metodo responsavel por atacar
func (person Person) Attack(attacked *Creature) {
	if !attacked.Defend() {
		attacked.Hp = attacked.Hp - rand.Intn(100)
	}

}

// Defend e metodo responsavel por defender
func (person Person) Defend() bool {
	if value := rand.Intn(1); value == 1 {
		return true
	}
	return false
}

// Creature struct responsavel pelas creatures
type Creature struct {
	Name string  `json:"name"`
	Exp  float64 `json:"experience"`
	Hp   int     `json:"hp"`
	Skill
}

// Attack metodo responsavel por atacar
func (creature Creature) Attack(attacked *Person) {
	if !attacked.Defend() {
		attacked.Hp = attacked.Hp - rand.Intn(100)
	}
}

// Defend metodo responsavel por defender a creature
func (creature Creature) Defend() bool {
	if value := rand.Intn(1); value == 1 {
		return true
	}
	return false
}

// Cave responsavel pelas caves
type Cave struct {
	Name      string     `json:"name"`
	Categorie string     `json:"categorie"`
	Creature  []Creature `json:"creature"`
}

// ListCave lista cave existentes
func (cave Cave) ListCave() []Cave {
	return []Cave{
		Cave{
			Name:      "Ice",
			Categorie: "Ice",
			Creature: []Creature{
				Creature{
					Name:  "Frozen",
					Exp:   10,
					Skill: Skill{Force: 10, Defense: 10},
					Hp:    120,
				},
				Creature{
					Name:  "ICE",
					Exp:   10,
					Skill: Skill{Force: 10, Defense: 5},
					Hp:    120,
				},
			},
		},
		Cave{
			Name:      "Fire",
			Categorie: "Fire",
			Creature: []Creature{
				Creature{
					Name:  "Dragon lord",
					Exp:   1000,
					Skill: Skill{Force: 100, Defense: 50},
					Hp:    1000,
				},
				Creature{
					Name:  "Fire lord",
					Exp:   100,
					Skill: Skill{Force: 70, Defense: 10},
					Hp:    120,
				},
			},
		},
		Cave{
			Name:      "Earth",
			Categorie: "Earth",
			Creature: []Creature{
				Creature{
					Name:  "Earthworm",
					Exp:   50,
					Skill: Skill{Force: 10, Defense: 1},
					Hp:    50,
				},
				Creature{
					Name:  "Mouse",
					Exp:   10,
					Skill: Skill{Force: 120, Defense: 70},
					Hp:    120,
				},
			},
		},
	}
}

func Fight(attack *Person, attacked *Creature) {
	for {
		attack.Attack(attacked)
		if attacked.Hp <= 0 {
			fmt.Println(attacked.Name, "está morto!")
			attack.Exp = attack.Exp + attacked.Exp
			break
		}

		attacked.Attack(attack)
		if attack.Hp <= 0 {
			fmt.Println(attack.Name, "está morto!")
			break
		}
	}
}
