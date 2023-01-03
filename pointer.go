package main

import "fmt"

type address struct {
	city, Provice, Coutry string
}

func changeAddress(ubah *address)  {
	ubah.city = "Yogyakarta"
}

func main() {
	var address1 address= address{"Jakarta", "Jakarta", "Indonesia"}
	var address2 *address = &address1

	address2.city = "Cilacap"

	*address2 = address{"malannnng","jawa Timur","Indonnesia"}
	fmt.Println(address1)
	fmt.Println(address2)

	alamat := address{"sleman","sleman","indonesia"}
	changeAddress(&alamat)
	fmt.Println(alamat)
}
