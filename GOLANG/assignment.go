package main 

import "fmt" 

func main(){ 

	var var1 = 5
	var var2 = 10  

	var2 += var1
	fmt.Println(var2)

	var2 -= var1
	fmt.Println(var2)

	var2 *= var1 
	fmt.Println(var2)

	var2 /= var1
	fmt.Println(var2)

	var2 %= var1
	fmt.Println(var2)
}
