package main 

import "fmt"

func main(){ 

		var working_day = "yes"
		var shop_open = "yes"
		var laptop_found = "yes"
		var high_ram = "no"
		var middle_ram = "yes"

		if(working_day == "yes"){
			if(shop_open == "yes"){ 
				}else{
					fmt.Println("Shop is not open")
				}
				if(laptop_found == "yes"){

				}else{
					fmt.Println("Laptop is not found")
				}
				if(high_ram == "yes"){
						fmt.Println("High Ram Laptop")
				}else if(middle_ram == "yes"){
						fmt.Println("Middle Ram Laptop")
					}else{
						fmt.Println("Low Ram Laptop")
					}
		}else{
			fmt.Println("Not a Holiday")
		}

}