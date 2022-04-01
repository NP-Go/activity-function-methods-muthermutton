package main

func main() {
	//Declare new item customer
	customer1 := customer{
		fName:    "Micheal",
		lName:    "Jordan",
		userName: "MJ2020",
		password: "1234567",
		email:    "MJ2020@gmail.com",
		phone:    12345678,
		address:  "18227 Capstan Greens Road Cornelius, NC 28031.",
	}
	//Insert Code here
	customer1.printAllInfo()
	customer1.getCredentials()
	customer1.getAddress()
}
