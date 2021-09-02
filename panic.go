package main

import "fmt"
//panic function untuk menghentikan program

func endApp() {
	fmt.Println("End App")
}

func runAPp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runAPp(true)
}