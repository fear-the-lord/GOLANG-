package main 

import "fmt" 

func main(){ 

	var grade = "C" 
	var marks = "60"

	switch marks{

	case "60" : 
		grade = "D"
		//fallthrough
	case "70" : 
		grade = "C"
		//fallthrough 
	case "80" : 
		grade = "B"
		//fallthrough
	case "90" :
		grade = "A"
		//fallthrough
	default : 
		grade = "O" 
	}

	switch grade{

	case "O" :
		fmt.Println("Outstanding")
	case "A" :
		fmt.Println("Very Good")
	case "B" :
		fmt.Println("Good")
	case "C" :
		fmt.Println("Fair")
	case "D" :
		fmt.Println("Fail")
	}
}