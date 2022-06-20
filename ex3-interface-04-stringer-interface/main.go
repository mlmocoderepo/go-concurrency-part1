package main

import "fmt"

type user struct {
	name  string
	email string
}

// Provide a custom function from the Stringer interface
// String implements the Stringer Interface to determine how the values are printed
func (u user) String() string {
	return fmt.Sprintf("%s <%s>", u.name, u.email)
}

func main() {

	u := user{
		name:  "John Doe",
		email: "johndoe@gmail.com",
	}

	u.name = "Sam Doe"

	fmt.Println(u)
}
