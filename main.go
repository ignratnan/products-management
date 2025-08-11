package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}

}
