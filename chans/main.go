package main

import "sync"

func main() {
	wg := &sync.WaitGroup{}

	test1(wg)

	wg.Wait()
}
