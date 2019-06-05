package main

import "fmt"

type hello map[string]int

func main() {
	world := hello{
		"ini": 10,
		"itu": 20,
	}

	fmt.Print(world["ini"] + world["itu"])
}
