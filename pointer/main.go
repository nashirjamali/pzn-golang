package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	// Contoh pass by value / tidak menggunakan pointer
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Sidoarjo"

	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference / menggunakan pointer
	address1_1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	address2_1 := &address1_1

	address2_1.City = "Bandung"

	fmt.Println(address1_1)
	fmt.Println(address2_1)

	// Operator *
	address1_2 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2_2 := &address1_2

	address2_2.City = "Bandung"

	*address2_2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1_2)
	fmt.Println(address2_2)

	// Menggunakan New()
	var address4 *Address = new(Address)
	fmt.Println(address4)

	// Pointer function
	var alamat = Address{
		City:     "Gresik",
		Province: "Jawa Timur",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
