package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["one"] = 1

	fmt.Println(m)
	fmt.Println(m["one"])
}