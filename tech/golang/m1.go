package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1337 svc")
	f, err := os.Create("testf")

	if err != nil {
		fmt.Println("Error")
	}
	f.Write([]byte("Test\n\x20\x00\x90"))
}
