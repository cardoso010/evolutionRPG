package models

// Action responsavel pelos metodos de acao
type Action interface {
	Attack(attacker interface{}, attacked interface{})
	Defend()
}

// Skill  struct responsavel skills
type Skill struct {
	Force   int `json:"force"`
	Defense int `json:"defense"`
}

// Person  struct responsavel pelos personagem
type Person struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Level      int     `json:"level"`
	Experience float64 `json:"experience"`
	Skill
	Hp int `json:"hp"`
}

// Attack metodo responsavel por atacar
func (person Person) Attack(attacker interface{}, attacked *interface{}) {

}

// Defend e metodo responsavel por defender
func (person Person) Defend() {

}

// Creature struct responsavel pelas creatures
type Creature struct {
	Name string `json:"name"`
	Skill
	Hp int `json:"hp"`
}

// Attack metodo responsavel por atacar
func (creature Creature) Attack(attacker interface{}, attacked *interface{}) {

}

// Defend metodo responsavel por defender a creature
func (creature *Creature) Defend() {

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
					Skill: Skill{Force: 10, Defense: 10},
					Hp:    120,
				},
				Creature{
					Name:  "ICE",
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
					Skill: Skill{Force: 100, Defense: 50},
					Hp:    1000,
				},
				Creature{
					Name:  "Fire lord",
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
					Name:  "earthworm",
					Skill: Skill{Force: 10, Defense: 1},
					Hp:    5,
				},
				Creature{
					Name:  "mouse",
					Skill: Skill{Force: 120, Defense: 70},
					Hp:    120,
				},
			},
		},
	}
}
