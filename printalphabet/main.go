package main

import (
	"fmt"
	"os"
)

func PrintRune(r rune) error {
	if _, err := fmt.Printf("%c", r); err != nil {
		return err
	}
	return nil
}

func main() {
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if err := PrintRune(r); err != nil {
			fmt.Fprintf(os.Stderr, "Error printing rune: %v\n", err)
			os.Exit(1)
		}
	}
}
