package main 

import ( 
"fmt"
"strings"
)

func main(){

	var var1 = strings.ToUpper("India") //Converts to Upper Case
	var var2 = strings.ToLower("iNDia") //Converts to Lower Case
	var var3 = strings.Title(" souvik ghosh") //Makes the initials in Upper Case
	var var4 = strings.TrimSpace(" Souvik Ghosh") //Trims space from First and Last
	var var5 = strings.Trim("India", "In") //Removes alphabets from the start and last
	var var6 = strings.Trim("India", "ia") //and not from the middle 
	var var7 = strings.TrimLeft("India", "In") //Removes matching strings only from the left 
	var var8 = strings.TrimRight("India", "dia") //Removes matching strings only from the right

	var var9 = strings.Count("india", "i") //counts the no:of times a given pattern is found in a string
	var var10 = strings.Count("india", "yia")

	var var11 = strings.Contains("India", "In") //checks whether a pattern is present in a string 
	var var12 = strings.Contains("India", "IND")//and returns the result in true or false
	var var13 = strings.ContainsAny("India", "Bci")//checks if any unit of a particular pattern is present in a string 
	var var14 = strings.ContainsAny("India", "bcf")//and returns the result in boolean

	var var15 = strings.Index("India", "a") //returns the first index of 'a'
	var var16 = strings.LastIndex("Indian", "n") //returns the last index of 'n'
	var var17 = strings.IndexAny("Indian", "nd")

	var var18 = strings.Replace("india", "i", ".", 1)//replaces one pattern with a given pattern in a string only in one position 
	var var19 = strings.Replace("india", "i", ".", 2)//two positions
	var var20 = strings.Replace("india", "i", ".", -1)//all positions

	var var21 = strings.Split("i.am.souvik", ".")//splits at "."
	var var22 = strings.SplitAfter("i.am.souvik", ".")//splits after "."

	var var23 = strings.Join([]string{"i","am","souvik"},".")//joins with "."
	var var24 = strings.Repeat("souvik",3)//repeats a string given no:of times

	fmt.Println(var1);
	fmt.Println(var2);
	fmt.Println(var3);
	fmt.Println(var4);
	fmt.Println(var5);
	fmt.Println(var6);
	fmt.Println(var7);
	fmt.Println(var8);
	fmt.Println(var9);
	fmt.Println(var10);
	fmt.Println(var11);
	fmt.Println(var12);
	fmt.Println(var13);
	fmt.Println(var14);
	fmt.Println(var15);
	fmt.Println(var16);
	fmt.Println(var17);
	fmt.Println(var18);
	fmt.Println(var19);
	fmt.Println(var20);
	fmt.Println(var21);
	fmt.Println(var22);
	fmt.Println(var23);
	fmt.Println(var24);
}