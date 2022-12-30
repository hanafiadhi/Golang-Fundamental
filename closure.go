package main

import "fmt"

func main() {
	name := "Hanafi"

	myname := func() {
		fmt.Println("this is function")
		//nama yang tadinya hanafi menjadi Adhi, jika tidak ingin berubah maka bisa di set menjadi variabel
		name := "Adhi"
		fmt.Println(name)
	}
	myname()
	fmt.Println(name)
}
