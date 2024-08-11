package main

import (
	"angry-embassies/config"
	"fmt"
)

func main() {
	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	document, err := deps.Client.InsertDocument("Hello, MongoDB!")
	if err != nil {
		fmt.Printf("deps.Client.InsertDocument error: %v\n", err)
		return
	}

	fmt.Println("Hello, GoLand!")
	fmt.Printf("Document inserted: %v\n", document)
}
