package main

import (
	"fmt"
	"golang-course/database"
)


func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}