package models

type Person struct {
	Name       string  `json:"name"`
	Experience float64 `json:"experience"`
	Level      int     `json:"level"`
}

type Creature struct {
	Name    string `json:"name"`
	Force   int    `json:"force"`
	Defense int    `json:"defense"`
}

type Cave struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
