package main

import "fmt"

func main() {
	
	var fnumber int
	fmt.Print("Enter Fist Number");
	fmt.Scan(&fnumber)

	var snumber int
	fmt.Print("Enter Seconds Number");
	fmt.Scan(&snumber);

	var total int = fnumber+ snumber

	fmt.Printf("Fist Number: %v Seconds Number: %v\n Total is: %v", fnumber, snumber, total)
	
}
