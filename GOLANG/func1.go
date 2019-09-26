package main 

import "fmt" 

func main(){

	var out = func1(8,9)
	var subOut, mulOut = func2(10, 2)
	fmt.Println(out)
	fmt.Println(subOut)
	fmt.Println(mulOut)

}

func func1(a, b int) int{

	//returning a single value in a function
	 var c = a + b
	return int(c)
}

func func2(a, b int) (int, int){

	//returning multiple values in a function
	var sub = a - b
	var mul = a * b 
	return sub, mul
}