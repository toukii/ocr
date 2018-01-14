package main

import (
	"fmt"

	"github.com/toukii/ocr/command-ocr"
)

func main() {
	if err := ocr.Command.Execute(); err != nil {
		fmt.Printf("%+v", err)
	}
}
