package main

import (
	"fmt"
	"os"
)

func getContent(style string) []byte {
	fileName := style+".txt"
	
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Style error")
		os.Exit(0)
	}
	return content
}
