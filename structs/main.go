package main

import "fmt"

type Node struct {
	int name `noder`
}

func main() {
	x := Node{}

	fmt.Println(x)
}
