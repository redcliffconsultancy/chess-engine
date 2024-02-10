package main

import (
	"fmt"
)

type Player struct {
	name          string
	rightHanded   bool
	maxServeSpeed int
}

func (p *Player) updateServeSpeed(speed int) {
	p.maxServeSpeed = speed
}

func main() {
	andrew := Player{name: "Bob", rightHanded: true, maxServeSpeed: 0}
	fmt.Println(andrew)
	andrew.updateServeSpeed(100)
	fmt.Println(andrew)
}
