package main

import (
	"fmt"
	"log"
	"os"

	"github.com/greymd/ojichat/generator"
)

func main() {
	name := "AAA"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	config := generator.Config{
		TargetName:        name,
		EmojiNum:          4,
		PunctiuationLebel: 1,
	}

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", result)
}
