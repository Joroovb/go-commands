package main

import (
	"fmt"
	"log"

	"github.com/joroovb/go-utils"
)

func main() {
	res, err := utils.Head("main.go", 5)
	if err != nil {
		log.Panic(err)
	}

	for _, line := range res {
		fmt.Println(line)
	}
}
