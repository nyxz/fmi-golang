// Драконов фрактал
package main

import "fmt"

type DragonFractal struct {
	currIndex int
	generation []int
}

var direction map[int]string = map[int]string {
	0 : "up",
	1 : "right",
	2 : "down",
	3 : "left",
}

// the translation of the directions
// after the rotation from one to another
var rotation map[int]int = map[int]int {
	0 : 3, // up -> left
	1 : 0, // right -> up
	2 : 1, // down -> right
	3 : 2, // left -> down
}

func (dragon *DragonFractal) Next() (rightDirection string) {
	if dragon.currIndex == len(dragon.generation) {
		dragon.populateNextGeneration()
	}
	rightDirection = direction[dragon.generation[dragon.currIndex]]
	dragon.currIndex++
	return
}

func (dragon *DragonFractal) populateNextGeneration() {
	genSize := len(dragon.generation)
	for revIdx := genSize - 1; revIdx >= 0 ; revIdx-- {
		dragon.generation = append(dragon.generation, rotateDirection(dragon.generation[revIdx]))
	}
}

func rotateDirection(direction int) int {
	return rotation[direction]	
}

func main() {
	dragon := DragonFractal{0, make([]int, 1, 1)} // for default up direction
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