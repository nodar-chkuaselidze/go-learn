package main

import (
	"fmt"
)

func main() {
	str := returnStr()
	fmt.Println(str)
	fmt.Println(" ")

	str2 := withPanic()
	fmt.Println(str2)
}

func returnStr() string {
	fmt.Println("starting")
	defer func() {
		fmt.Println("Hello DEFER")
	}()

	return func() string {
		fmt.Println("returning..")
		return "Hello World"
	}()
}

func withPanic() string {
	fmt.Println("starting with panic")
	defer func() {
		r := recover()

		fmt.Println("recovered.. : ", r)
	}()

	return func() string {
		panic("WWWWAAAA")
		return "Hello from withPanic"
	}()
}
