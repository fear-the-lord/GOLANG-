package main 

include "fmt"

func main(){

	var ptr *int 
	ptr = 23
	fmt.Println(&ptr)
	fmt.Println(ptr)
}