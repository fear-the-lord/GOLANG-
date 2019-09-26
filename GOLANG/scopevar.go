package main 

import "fmt"

//var1 is a global variable 
var var1 = "Hello From outside main"

func main(){

		fmt.Println(var1)
		var var2 = "Hello from inside main"
		//var2 is a local variable and cannot be accessed from outside the function
		fmt.Println(var2)
		func1()

}

func func1(){

	var var3 = "Hello from inside func1"
	fmt.Println(var3)
	fmt.Println(var1)
}