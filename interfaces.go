package main

import (
	"fmt"
)


type animal interface {
	move() string
}


type dog struct{}

type fish struct{}

type bird struct{}


func (dog) move() string {
	return "I'm a dog and do walking"
}

func (fish) move() string {
	return "I'm a fish and do swimming"
}

func (bird) move() string {
	return "I'm a bird and do flying"
}


func moveAnimal(a animal) {
	fmt.Println(a.move())
}


func main() {
	p := dog{}
	f := fish{}
	b := bird{}

	moveAnimal(p)
	moveAnimal(f)
	moveAnimal(b)

}