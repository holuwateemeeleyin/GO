package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Identifying data types can be done in 2 ways:
	//1. By using %T format specifier, and
	//2. reflect.TypeOf function from the reflect package

	//Using %T
	var scores int = 50
	var name string = "Timi"
	var isSuccess bool = true
	var amount float32 = 170.35

	fmt.Printf("The variable score %v is %T \n", scores, scores)
	fmt.Printf("The variable name %v is %T \n", name, name)
	fmt.Printf("The variabe isSuccess %v is %T \n", isSuccess, isSuccess)
	fmt.Printf("The variable amount %v is %T \n", amount, amount)

	//Using reflect.TypeOf
	fmt.Println("Using reflect.Typeof: ")
	var school string = "Banwo"
	var grade int = 40

	fmt.Printf("The variable school %v is of datatype %v \n", school, reflect.TypeOf(school))
	fmt.Printf("The variable grade %v is of datatype %v \n", grade, reflect.TypeOf(grade))
	fmt.Println("The variable school", school, "is of datatype", reflect.TypeOf(school))
}
