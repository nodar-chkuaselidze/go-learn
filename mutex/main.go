package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	someMu := &sync.Mutex{}

	fmt.Println(someMu)

	go func() {
		someMu.Lock()
		fmt.Println("lock all stuff")
	}()

	go func() {
		<-time.NewTimer(time.Second * 2).C
		someMu.Unlock()
	}()

	go func() {
		someMu.Lock()
		fmt.Println("just locking this mutex")
		someMu.Unlock()
	}()

	go func() {
		someMu.Lock()
		fmt.Println("unlocked? ")
		someMu.Unlock()
	}()

	<-time.NewTimer(time.Second * 4).C
}
