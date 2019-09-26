package main 

import "fmt"

func main(){

	var multiArray = [3][2] int{{1, 2}, {3, 4}, {5, 6}}
	var multiArray2 = [...][3] int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}} 

	fmt.Println(multiArray)
	fmt.Println(multiArray2)
	fmt.Println(multiArray[2][0])
}