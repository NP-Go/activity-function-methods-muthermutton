package main

import "fmt"

//Declare a struct
type customer struct {
	fName    string
	lName    string
	userName string
	password string
	email    string
	phone    int
	address  string
}

//Create Methods
func (c customer) getCredentials() (string, string) {
	fmt.Println(c.userName, c.password)
	return c.userName, c.password
}

func (c customer) getAddress() string {
	fmt.Println(c.address)
	return c.address
}

func (c customer) printAllInfo() {
	fmt.Println(c.fName, c.lName, c.userName, c.password, c.email, c.phone, c.address)
}
