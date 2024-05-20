package main

import (
	"fmt"
	"os"
)

func main() {
	arg1 := os.Args[1]
	fmt.Println(arg1)
	data, err := os.Open(arg1)

	if err !=nil{
		panic(err)
	} 
	b1 := make([]byte, 5)
    n1, err := data.Read(b1)
    if err!=nil{
		panic(err)
	}
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
}