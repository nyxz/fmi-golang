package main

import "fmt"

type DragonFractal struct {
	currStep int
	rotationPoint int
	generation []int
}

var direction map[int]string = map[int]string {
	0 : "up",
	1 : "right",
	2 : "down",
	3 : "left",
}

var rotation map[int]int = map[int]int {
	0 : 3, // up -> left
	1 : 0, // right -> up
	2 : 1, // down -> right
	3 : 2, // left -> down
}

func (dragon *DragonFractal) Next() (rightDirection string) {
	dragon.updateRotationPoint()
	rotatedPoint := dragon.getRotatedPoint()

	dragon.generation = append(dragon.generation, rotatedPoint)
	rightDirection = direction[rotatedPoint]
	dragon.currStep++
	
	return
}

func (dragon *DragonFractal) getRotatedPoint() (rotatedPoint int) {
	diff := dragon.currStep - dragon.rotationPoint
	mirrorPoint := dragon.generation[dragon.rotationPoint - diff]
	rotatedPoint = rotation[mirrorPoint]
	return
}

func (dragon *DragonFractal) updateRotationPoint() {
	if isPowerOf2(dragon.currStep + 1) {
		dragon.rotationPoint = dragon.currStep
	}
}

func isPowerOf2(x int) bool {
	return (x != 0) && (x & (x - 1) == 0)
} 

func rotateDirection(direction int) int {
	return rotation[direction]	
}

func main() {
	// len = 1 for default UP direction
	dragon := DragonFractal{0, 0, make([]int, 1, 1)}
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