package main

import (
	"strings"
	"fmt"
)

var SEPARATOR string = "/"

func parsePath(complexPath string) (simplePath string) {
	dirs := []string{}
	for _, dir := range strings.Split(complexPath, SEPARATOR) {
		if isBackwardSymbol(dir) {
			dirs = dirs[:(len(dirs) - 1)]
		} else if isBlank(dir) {
			// ignore
		} else {
			dirs = append(dirs, dir)
		}
	}

	simplePath = buildSimplePath(dirs)

	return
}

func buildSimplePath(dirs []string) string {
	if len(dirs) == 0 {
		return SEPARATOR
	}
	simplePath := strings.Join(dirs, SEPARATOR)
	return SEPARATOR + simplePath + SEPARATOR
}

func isBackwardSymbol(dir string) bool {
	return dir == ".."
}

func isBlank(dir string) bool {
	return dir == "" || dir == "." 
}

func main() {
	// fmt.Println(parsePath("/home/maistora/test/dir/../../anotherTest"))
	fmt.Println(parsePath("/home/maistora/test/dir/../../anotherTest/../../"))
}