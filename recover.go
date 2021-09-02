package main

import "fmt"

// untuk menangkap error panic
func endApp() {
	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Message errpr ", message)
	}
	fmt.Println("Aplikasi Selesai")
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