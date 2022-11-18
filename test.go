package main

import "fmt"

func main () {
	var name string
	var names = []string {}
	for {
	fmt.Printf("Please Enter your name %v\n",name)
	fmt.Scan(&name)
	fmt.Printf("Welcome %v \n",name)
	names = append(names, name)
	fmt.Printf("All names %v\n",names)
	}
	
}