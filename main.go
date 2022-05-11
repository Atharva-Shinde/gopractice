// package main

// import (
// 	"fmt"
// )

// func generateNumbers(total int) {
// 	for i :=1; i<= total; i++{
// 		fmt.Printf("Generating number %d\n", i)
// 	}
// }

// func printNumbers(){
// 	for i :=1; i<= 3; i++{
//                 fmt.Printf("Printing numbers")
//         }
// }

// func main(){
// 	printNumbers();
// 	generateNumbers(3);

// }

package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printNumbers(&wg)
	go generateNumbers(3, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
