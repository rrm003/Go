package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

var (
	info   []string
	system map[string]string
)

func main() {
	args := os.Args[:]
	if len(args) < 2 {
		fmt.Println("Use cmd : hello help ")
		os.Exit(0)
	}

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
	temp := args[1:]
	for inp := range temp[:] {

		if temp[inp] == "help" {
			fmt.Println("Try hello <input>\nInput :  ")
			for i := range system {
				fmt.Println("\thello ", i)
			}
			os.Exit(0)
		}
		if system[temp[inp]] == "" {
			fmt.Println("Invalid Command\nTry hello help")
			os.Exit(0)
		}
		fmt.Println(temp[inp], " : ", system[temp[inp]])
		//fmt.Println(temp[inp])
	}
	//fmt.Println("...exiting program!")
}
