package main

import "fmt"

type User struct {
	ID            int     `json:"id"`
	Firstname     string  `json:"first_name"`
	Lastname      string  `json:"last_name"`
	Age           int     `json:"age"`
	DriverLicense *string `json:"driver_license"`
}

func (u User) GetFullName() string {
	return u.Firstname + " " + u.Lastname
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\nFirstname: %s\nLastname: %s\nAge: %d", u.ID, u.Firstname, u.Lastname, u.Age)
}

func main() {
	user := new(User)
	user.ID = 1
	user.Firstname = "Joshua"
	user.Lastname = "Suherman"
	user.Age = 25

	user2 := User{
		ID:        2,
		Firstname: "Joshua",
		Lastname:  "Suherman",
		Age:       25,
	}

	driverLicense := "A1234567"

	user.DriverLicense = &driverLicense

	driverLicense = "B7654321"

	user2.DriverLicense = &driverLicense

	// variable by value (copy)
	// variable by reference (pointer)

	// driverLicense = "B7654321"

	fmt.Println(*user.DriverLicense)
	fmt.Println(*user2.DriverLicense)

	fmt.Println(user.GetFullName())

	fmt.Printf("%T\n", user.GetFullName())

}
