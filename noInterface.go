package main

import (
	"fmt"
)


type dog struct{}

type fish struct{}

type bird struct{}


func (dog) walk() string {
	return "I'm a dog and do walking"
}

func (fish) swim() string {
	return "I'm a fish and do swimming"
}

func (bird) fly() string {
	return "I'm a bird and do flying"
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	p := dog{}
	f := fish{}
	b := bird{}

	moveDog(p)
	moveFish(f)
	moveBird(b)

}