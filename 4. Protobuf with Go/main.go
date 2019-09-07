package main

import (
	"fmt"

	simple "./proto/go"
)

func main() {
	doSimple()
}

func doSimple() {
	s := simple.Simple{
		Id:     10,
		IsGood: true,
		Name:   "Gaurav",
		Score:  []int32{10, 20},
	}

	fmt.Println(s)
}
