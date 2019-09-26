package main 

import "fmt"

func main(){

	var var1 = 3
	fmt.Println(var1)

	var var2 int 
	var2 = 9
	fmt.Println(var2)

	var var3 float64 = 3.89 
	fmt.Println(var3)
	fmt.Printf("var3 type is %T\n", var3)

	var var4, var5, var6 int
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)

	 var var7, var8, var9 = "foo", 8, "bar" 
	 fmt.Println(var7)
	 fmt.Println(var8)
	 fmt.Println(var9)

	 var x = 10 
	 fmt.Println(x)
	 x = 8 
	 fmt.Println(x)

	 const y = 12 
	 fmt.Println(y)
	 /*y = 18 
	 fmt.Println(y) cannot assign to y*/

}