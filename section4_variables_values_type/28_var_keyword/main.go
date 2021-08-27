package main

import "fmt"

//x := 22 does not compile; syntax error: non-declaration statement outside function body

// 'var' might be useful for constants used in the whole file
// the scope of my_var_with_large_scope is package level
var my_var_with_large_scope = 65

/*
You can declare type for var, variable "count" is of type int and default value is 0

If no explicit initialization is provided, the variable or value is given a default value.
Each element of such a variable or value is set to the zero value for its type:
false for booleans, 0 for numeric types, "" for strings,
and nil for pointers, functions, interfaces, slices, channels, and maps.
*/
var count int
func main() {
	/*
	Use short declaration operator ':=' inside function body
	Best practice: limit the scope of variables by declaring them inside function body with ':=' as much as possible
	*/
	x := 45
	fmt.Println("I cannot use short declaration operator ':=' to declare variable outside of function body, x = ", x)

	var y = 34
	fmt.Println(y)

	fmt.Println("I can declare variable outside of a function body as 'var', var my_var_with_large_scope = ", my_var_with_large_scope)

	fmt.Println("Example of int var with default value, count = ", count)
}
