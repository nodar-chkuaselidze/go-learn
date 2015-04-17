package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	Locked bool
	Locker chan bool
}

func (m *Mutex) Lock() {
	if m.Locker == nil {
		m.Locker = make(chan bool, 1)
	}

	if m.Locked {
		<-m.Locker
		m.Locked = false
	} else {
		m.Locked = true
	}
}

func (m *Mutex) Unlock() {
	m.Locker <- true
}

func main() {
	mu := &Mutex{}

	fmt.Println(mu)

	go func() {
		mu.Lock()
		fmt.Println("locking ...")
	}()

	go func() {
		mu.Lock()
		fmt.Println("unlocked !!")
	}()

	go func() {
		<-time.NewTimer(time.Second * 2).C
		mu.Unlock()
		fmt.Println("unlocking..")
	}()

	<-time.NewTimer(time.Second * 4).C
	fmt.Println("finished..")
}
