package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jp := person{
		firstName: "Juan",
		lastName:  "Giraldo",
		contact: contactInfo{
			email:   "code@qjuanp.dev",
			zipCode: 1234, // If there is a multiline value declaration, the last line should have ','
		},
	}

	fmt.Printf("%+v", jp)

	jp.firstName = "another"
	jp.lastName = "person"

	fmt.Printf("%+v", jp)
}
