package main

import "fmt"

func endApp() {
	if message := recover(); message != nil {
		fmt.Println("Error message : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan....")
}
func main() {
	runApp(true)
}
