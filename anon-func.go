package main

import "fmt"

//intinya anonymouse function adalah memasukan function ke variabel atau memasukan function jadi parameter
type Blacklist func(string) bool

func registerUser(name string,  blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("You are blocked",name)
	} else {
		fmt.Println("Welcome", name)
	}
}


func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing" , blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
