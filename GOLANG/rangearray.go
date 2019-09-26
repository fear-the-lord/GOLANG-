package main 

import "fmt"

func main() {

	var array1 = [5]int {1,2,3,4,5}
	fmt.Println(array1)

	for c := range array1 { 
		fmt.Println(c) // displays the index of an array 
	}

	var array2 = [6]int {9, 8, 0, 7, 6, 2}

	for c := range array2 { 
		fmt.Println("The element at index", c , "is" , array2[c])
	}

	var slice = [] int{10, 11, 12, 13, 14}

	for c := range slice{ 
		fmt.Println("[Slice] The element at index", c, "is", slice[c])
	}

	var map1 = map[string]string{"Japan":"Tokyo", "India":"New Delhi", "China":"Beijing", "Russia":"Moscow"}

	for c := range map1{
		fmt.Println("The capital of", c , "is", map1[c])
	}

	var map2 = map[int]int{2:4, 3:9, 4:16, 5:25}

	for c,r := range map2{
		fmt.Println("The square of",c , "is", r)
	}
}