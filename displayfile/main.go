package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length > 2 {
		fmt.Println("Too many arguments")
	} else if length == 1 {
		fmt.Println("File name missing")
	} else if length == 2 {
		file, _ := os.Open(os.Args[1])
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}

// test
