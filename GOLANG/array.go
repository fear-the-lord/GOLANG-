package main 

import "fmt"

func main(){

	var strings = [3] string{"Hello","World","Program"}
	var quotes = [4] string{"Good", "Morning"}//2 spaces will be added since the array is of size 4 but there are only 2 elements 
	var quotes1 = [...] string{"Good", "Evening"}//Dynamic

	var num = [3] int{1,2,3}
	var num1 = [4] int{4, 9, 8}

	var num2 = [3] int{}
		num2[0] = 1
		num2[1] = 5
		num2[2] = 9

	fmt.Println(strings)
	fmt.Println(quotes)
	fmt.Println(strings[0])
	fmt.Println(quotes1)

	fmt.Println(num)
	fmt.Println(num[1] + num1[0])
	fmt.Println(num1)//one 0 for the size of the array is 4 but there are 3 elements
	fmt.Println(num2)

	fmt.Println(len(quotes1))
	fmt.Println(len(num1))
	fmt.Println(len(num2))
	fmt.Println(len(strings))

}