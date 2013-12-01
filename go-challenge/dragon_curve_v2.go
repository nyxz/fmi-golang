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

const DEFAULT_DIRECTION int = 0 // up

func (dragon *DragonFractal) Next() (rightDirection string) {
	if dragon.shouldInit() {
		rightDirection = dragon.init()
		return
	}
	dragon.updateRotationPoint()
	rotatedPoint := dragon.getRotatedPoint()

	dragon.generation = append(dragon.generation, rotatedPoint)
	rightDirection = direction[rotatedPoint]
	dragon.currStep++
	
	return
}

func (dragon *DragonFractal) updateRotationPoint() {
	if isPowerOf2(dragon.currStep + 1) {
		dragon.rotationPoint = dragon.currStep
	}
}

func (dragon *DragonFractal) getRotatedPoint() (rotatedPoint int) {
	diff := dragon.currStep - dragon.rotationPoint
	mirrorPoint := dragon.generation[dragon.rotationPoint - diff]
	rotatedPoint = rotation[mirrorPoint]
	return
}

func (dragon *DragonFractal) shouldInit() bool {
	return dragon.generation == nil
}

func (dragon *DragonFractal) init() (initDirection string) {
	dragon.generation = make([]int, 1, 1)
	dragon.generation[0] = DEFAULT_DIRECTION
	initDirection = direction[DEFAULT_DIRECTION]
	return
}

func isPowerOf2(x int) bool {
	return (x != 0) && (x & (x - 1) == 0)
} 

func rotateDirection(direction int) int {
	return rotation[direction]	
}

func main() {
	dragon := DragonFractal{
		currStep : 0,
		rotationPoint : 0,
		generation : nil}
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