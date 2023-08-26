package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	Print(ScriptBanner())

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	Print("Scanning for lanugage...")

	dir_content, err := os.ReadDir(os.Args[1])

	if err != nil {
		Print("Failed to parse directory!")
		os.Exit(1)
	}

	var language string = ""

	for _, item := range dir_content {
		switch item.Name() {
		case "main.py":
			language = "python"
		case "main.go":
			language = "golang"
		}
	}

	Print("Creating Dopamine DevTools with preset: " + strings.ToUpper(language) + "...")
}
