package main 

import (
	"fmt"
	"math"
)

func main() {

	var var1 = math.Ceil(-1.90)

	var var2 = math.Sqrt(36)

	var var3 = math.Remainder(3, 9)

	var var4 = math.Max(-5, -9)

	var var5 = math.Min(-5, -9)

	var var6 = math.Mod(3, 9)
	
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)

}