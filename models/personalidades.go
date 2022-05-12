package models

type Personalidade struct {
	Id       int    `json: "Id"`
	Nome     string `json:"nome"`
	Historia string `json:"história"`
}

var Personalidades []Personalidade
