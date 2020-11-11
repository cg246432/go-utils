package main

import (
	"fmt"
	"sync"
)

func regularForLoop(inputSlice []string){

	for i:=0; i < len(inputSlice); i++{
		fmt.Println("Running regular for loop.")
		val := inputSlice[i]
		fmt.Println(val)
	}
}

func concurrentForLoop(inputSlice []string){
	// Only good when members of for loop don't depend on one another
	var wg sync.WaitGroup
	// Tell the 'wg' WaitGroup how many threads/goroutines
	//   that are about to run concurrently.
	wg.Add(len(inputSlice))
	fmt.Println("Running concurrent for loop.")
	for i := 0; i < len(inputSlice); i++ {
		// Spawn a thread for each iteration in the loop.
		// Pass 'i' into the goroutine's function
		//   in order to make sure each goroutine
		//   uses a different value for 'i'.
		go func(i int) {
			// At the end of the goroutine, tell the WaitGroup
			//   that another thread has completed.
			defer wg.Done()
			val := inputSlice[i]
			fmt.Printf("i: %v, val: %v\n", i, val)
		}(i)
	}
	// Wait for `wg.Done()` to be exectued the number of times
	//   specified in the `wg.Add()` call.
	// `wg.Done()` should be called the exact number of times
	//   that was specified in `wg.Add()`.
	// If the numbers do not match, `wg.Wait()` will either
	//   hang infinitely or throw a panic error.
	wg.Wait()



}




func main(){
	slice := []string{"a", "b", "c", "d", "e"}

	regularForLoop(slice)
	concurrentForLoop(slice)

}