package models

type Personalidade struct {
	Id       int    `json: "Id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
