package main

// defer function adalah function yang bisa kita jadwalkan untuk diekseskusi setelah sebuah function selesai diekseskusi
// walaupun ada error tetap dieksekusi

import "Fmt"

func logging() {
	fmt.Println("Selesai memanggil function s")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}