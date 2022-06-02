package main

import "fmt"
/*
var (
	name     string
	age      int
	location string
)

var (
	name, location  string
	age             int
)
*/
var (
	name     = "Brakula"
	age      =  399
	location = "England"
)



func main() {
	//name, location := "Prince Oberyn", "Dorne"
	//age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)
}