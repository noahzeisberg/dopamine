package main

import "fmt"

func PrintR(msg any) {
	fmt.Print(msg)
}

func Print(msg ...any) {
	fmt.Println(msg...)
}
