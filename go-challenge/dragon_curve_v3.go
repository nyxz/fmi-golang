package main

import "fmt"

const DEFAULT_DIRECTION string = "up"

type DragonFractal struct {
	currStep int
	rotationPoint int
	generation []string
}

var rotation = map[string]string {
	"up" : "left",
	"right" : "up",
	"down" : "right",
	"left" : "down",
}

func (dragon *DragonFractal) Next() string {
	if dragon.generation == nil {
		dragon.generation = []string { DEFAULT_DIRECTION }
		return dragon.generation[dragon.currStep]
	}
	if isPowerOf2(dragon.currStep + 1) {
		dragon.rotationPoint = dragon.currStep
	}

	diff := dragon.currStep - dragon.rotationPoint
	mirrorPoint := dragon.generation[dragon.rotationPoint - diff]
	rotatedPoint := rotation[mirrorPoint]

	dragon.generation = append(dragon.generation, rotatedPoint)
	dragon.currStep++
	return dragon.generation[dragon.currStep]
}

func isPowerOf2(x int) bool {
	return (x != 0) && (x & (x - 1) == 0)
} 

func main() {
	dragon := new(DragonFractal)
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
}