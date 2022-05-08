package main

import (
	"fmt"
)

func generateNumbers(total int) {
	for i :=1; i<= total; i++{
		fmt.Printf("Generating number %d\n", i)
	}
}

func printNumbers(){
	for i :=1; i<= 3; i++{ 
                fmt.Printf("Printing numbers")
        }
}

func main(){
	printNumbers();
	generateNumbers(3);
	
}
