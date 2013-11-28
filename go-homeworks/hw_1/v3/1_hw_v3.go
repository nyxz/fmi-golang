package main

import (
	"fmt"
	"path"
)

func parsePath(complexPath string) string {
	simplePath := path.Join("/", complexPath)
	if len(simplePath) != 1 {
		return simplePath + "/"
	}
	return simplePath
}

func main() {
	fmt.Println(parsePath("home/maistora/./../pehso/Desktop/../.."))
}