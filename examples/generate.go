package main

import (
	"fmt"
	"uwuid"
)

func main() {
	id, err := uwuid.NewUwUID()
	if err != nil {
		panic("aww ney, something went ^w^ wwong, we *notices buldge* c-couwd'nt genyewate a UwUID :'(")
	}
	fmt.Printf("Hewe is youw UwUID Sempai : %v\n", id)
}
