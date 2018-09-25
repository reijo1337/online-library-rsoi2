package main

import (
	"fmt"
)

func main() {
	_, err := SetUpDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
}
