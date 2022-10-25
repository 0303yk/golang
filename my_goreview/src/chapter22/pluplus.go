// waitgroup
package main

import (
	f "fmt"
	s "sync"
	t "time"
)
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.

func process(i int , wg *s.WaitGroup){
	f.Println("started Goroutine ", i)
	t.Sleep(2 * t.Second)
	f.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg s.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	f.Println("All go routines finished executing")
}