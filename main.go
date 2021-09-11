package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var scanner *bufio.Scanner
var errOutp error
var command string

func main() {
	scanner = bufio.NewScanner(os.Stdin)

 	//clrScrn()

	fmt.Printf("Welcome to %s, the next generation of %s!\n", color.Magenta("GoCMD"), dispColor("NetCMD", Teal))

	for true {
		path, err := os.Getwd()
		errOutp = err
		fmt.Printf("%s",path)
		command = question("\n" + dispColor("$ ",DarkGray))
	}
}

func clrScrn() {
	fmt.Print("\033[H\033[2J")
}

func dispColor(text string, color string) string {
	return color + text + DefaultColor
}

func question(text string) string {
	fmt.Printf(text)
	scanner.Scan()
	return scanner.Text()
}

const (
	Black        = "\033[1;30m"
	Red          = "\033[1;31m"
	Green        = "\033[1;32m"
	Yellow       = "\033[1;33m"
	Purple       = "\033[1;34m"
	Magenta      = "\033[1;35m"
	Teal         = "\033[1;36m"
	White        = "\033[1;37m"
	DarkGray	 = "\033[1;90m"
	DefaultColor = "\033[0m"
)
