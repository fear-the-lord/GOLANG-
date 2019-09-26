package main 

import "fmt" 

func func1(){
	fmt.Println("Hello from func1")
}

func func2(person string){
		fmt.Println("Hello " + person + " from func2")
}

func func3(person string, age string){
 		fmt.Println("Hello " + person + " from func3 whose age is " + age)
}

func func4(a, b int){

fmt.Println("Hello from func4")
	fmt.Println(a + b)
}

func main(){
	fmt.Println("Hello from main")
	func1()
	func2("Souvik")
	func3("Souvik", "21")
	func4(8,9)
}