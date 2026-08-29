package main

type Address struct {
	City string
}

func (address Address) Location() string {
	return address.City
}

type Person struct {
	Name string
	Address
}
