package main

import (
	"fmt"
	"strings"
)

var SPLITTER string = "/"
var BACKWARD_DIR string = ".."

/* Stack implementation */
type Element struct {
	value interface{}
	next *Element
}

type Stack struct {
	top *Element
	size int
}

func (s *Stack) len() int {
	return s.size
}

func (s *Stack) push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) pop() *Element {
	if s.size == 0 {
		return nil
	}
	popped := s.top
	
	s.size -= 1
	s.top = s.top.next 
	
	return popped
}

/* Path manipulation */
type Pather struct {
	value string
	stack *Stack
}

func (p *Pather) getPath() string {
	if p.stack == nil || p.stack.len() == 0 {
		return SPLITTER
	}
	for elem := p.stack.pop(); elem != nil; {
		p.value = SPLITTER + elem.value.(string) + p.value 
		elem = p.stack.pop()
	}
	p.value += SPLITTER
	return p.value
}

func (p *Pather) addDir(value string) {
	if value == "" {
		return
	} else if isBackwardDir(value) {
		p.stack.pop()
	} else {
		p.stack.push(value)
	}
}

/* Task specific methods */
func isBackwardDir(value string) bool {
	if value == BACKWARD_DIR {
		return true
	}
	return false
}

func parsePath(initialPath string) string {
	dirs := strings.Split(initialPath, SPLITTER)
	path := new(Pather)
	path.stack = new(Stack)

	for i := 0; i < len(dirs); i++ {
		path.addDir(dirs[i])
	}
	return path.getPath()
}

func main() {
	// val := parsePath("/home/maistora/Dev/Workspace/../../Pesho")
	val := parsePath("home/maistora/../Dev/../Workspace/krasta/gavra/../../Pesho/../")
	fmt.Println(val)
}