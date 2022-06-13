package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]
	dat, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	io.Copy(os.Stdout, dat)

}
