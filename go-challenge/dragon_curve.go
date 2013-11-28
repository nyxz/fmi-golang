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
	fmt.Println("index: ", dragon.currIndex, ", generation len: ", len(dragon.generation))
	if dragon.currIndex == len(dragon.generation) {
		fmt.Println(dragon.generation)
		dragon.populateNextGeneration()
	}
	rightDirection = direction[dragon.currIndex]
	dragon.currIndex++
	return
}

func (dragon *DragonFractal) populateNextGeneration() {
	fmt.Println("Populating new generation...")
	oldGen := dragon.generation
	dragon.generation = make([]int, len(oldGen), 2 * cap(oldGen))
    copy(dragon.generation, oldGen)

	fmt.Println(oldGen)
	fmt.Println(dragon.generation)
	bias := len(dragon.generation)
	for i := 0; i < len(oldGen); i++ {
		dragon.generation[i + bias] = rotateDirection(oldGen[i])
	}
}

func rotateDirection(direction int) int {
	return rotation[direction]	
}

func main() {
	make([]int, 0, 1)
	dragon := DragonFractal{0, make([]int, 0, 1)} 
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
	fmt.Println(dragon.Next())
}