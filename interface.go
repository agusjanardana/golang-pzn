package main

//tipe data abstract, tidak ada interface langsung
// sebuah interface berisikan difinisi method
// biasanya interface sebagai kontrak

import "fmt"

type HasName interface{
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello" , hasName.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main( ) {
	var agus Person
	agus.Name = "Agus"

	SayHello(agus)
}