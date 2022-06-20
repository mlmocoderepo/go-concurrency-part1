package main

import "fmt"

func checkType(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("Value is of type int")
	case string:
		fmt.Println("Value is of type string")
	case bool:
		fmt.Println("Value is of type bool")
	case CustomStr:
		fmt.Println("Value is of type CustomStr")
	}
}

type CustomInterface interface{}

type CustomStr string

func (c CustomStr) String() string {
	return c.String()
}

func main() {

	// CustomInterface accepts 'any' form of data type. In this case, it accepts a type of CustomStr
	var i CustomInterface = CustomStr("James")

	if _, ok := i.(CustomStr); ok {
		fmt.Printf("Type of i is an %s: %v\n", fmt.Sprintf("%T", i), ok)
	}

	// CustomInterface accepts 'any' form of data type. In this case, it accepts a type of CustomStr
	var ci CustomInterface = 1

	if _, ok := ci.(CustomInterface); ok {
		fmt.Printf("Type of ci is an %s: %v\n", fmt.Sprintf("%T", ci), ok)
	}

	c := CustomStr("Jack")
	checkType(c)

	s := "ten"
	checkType(s)

	b := true
	checkType(b)

	// Type ASSERTION here WILL NOT work
	// if _, ok := c.(CustomStr); ok {
	// 	fmt.Println("c is of custom type")
	// }

	// fmt.Printf("%T", c)
}
