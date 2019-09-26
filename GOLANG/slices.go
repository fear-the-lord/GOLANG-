package main

import "fmt"

func main(){

	var slice1 = [] int{1,2,3,4,5}
	fmt.Println(slice1)
	var slice2 = [] int{}
	fmt.Println(slice2)
	var slice3 = make([]int,4,4)//assigning size to a slice 
	fmt.Println(slice3)
	slice1 = append(slice1,6,7)//adding values to a slice
	fmt.Println(slice1)
	slice1 = append(slice1[:2], slice1[5:]...)//to delete all elements from position 2 upto position 5 
	fmt.Println(slice1)
	fmt.Println(len(slice1))//length
	fmt.Println(len(slice2))
	fmt.Println(cap(slice1))//capacity
	fmt.Println(cap(slice2))
}