package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//variable definition
	//var variable_list optional_data_type
	//var x, y, z int
	//var a, b, c string
	//var t1 = 12
	var t1 = 112
	fmt.Println(t1)

	//static type
	var v1 int = 12
	fmt.Println(v1)

	//Dynamic type
	v2 := 23
	fmt.Println(v2)

	//mixed variable declaration
	var v3, v4, v5 = 1, 2, "string"
	fmt.Printf("a : %T b: %T C: %T \n", v3, v4, v5)

	//constant
	const cnt string = "value of "
 
}
