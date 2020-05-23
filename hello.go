package main

import (
	"strings"
	"syscall"
)

var (
	info   []string
	system map[string]string
)

func main() {
	info = syscall.Environ()
	var key string
	system := make(map[string]string)
	for i := range info {
		rslt := strings.Split(info[i], "=")
		if rslt[0] != "" || rslt[1] != "" {
			key = rslt[0]
			system[key] = rslt[1]
		}
	}
}
