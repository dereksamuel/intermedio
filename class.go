package main

import "fmt"

type abilities struct {
	strong float32
	speed  float32
	hax    float32
}

type Animal struct {
	uuid      string
	name      string
	abilities abilities
}

func Class() {
	a := Animal{}
	fmt.Printf("%v\n", a)

	a.uuid = "axr8a4sa5-a4sa5a512xc2bf5d6a"
	a.name = "Lion"
	a.abilities.speed = 12.8
	a.abilities.strong = 82.8
	a.abilities.hax = 49

	fmt.Printf("%v\n", a)
}
