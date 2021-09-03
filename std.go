package main

import (
	"container/list"
	"fmt"
)

func main(){
	data := list.New()
	data.PushBack("Agus")
	data.PushBack("Janardana")
	data.PushBack("Abasan")

	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
