package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	theMajor()
}

func theMajor() {

	input := takeInput()

	if input == "" {
		fmt.Println("EMPTY")
		theMajor()
	}

	Op := SelectOp()

	switch Op {
	case "a":
		result1, er1 := hex(input)
		if er1 != nil {
			fmt.Println(input, "is not valid hex.")
			theMajor()
		} else {
			fmt.Println(input, "in HEX to DEC is: ", result1)
			theMajor()
		}
	case "b":
		result2, er2 := bin(input)
		if er2 != nil {
			fmt.Println(input, "is not valid binary.")
			theMajor()
		} else {
			fmt.Println(input, "in BIN to DEC is: ", result2)
			theMajor()
		}
	case "c":
		result3, result4 := dec(input)
		fmt.Println(input, "in DEC to BIN is: ", result3)
		fmt.Println(input, "in DEC to HEX is: ", result4)
		theMajor()
	}
}

func takeInput() string {
	fmt.Println("Enter a STRING to Conver or q to QUIT")

	var theNum string
	fmt.Scan(&theNum)

	if theNum == "q" {
		os.Exit(1)
	}

	return theNum
}

func SelectOp() string {

	fmt.Println("Select an OPERATION  or q to QUIT")
	fmt.Println("a. for HEX to DEC")
	fmt.Println("b. for BIN to DEC")
	fmt.Println("c. for DEC to HEX & BIN")

	var theOperation string
	fmt.Scan(&theOperation)

	if theOperation == "q" {
		os.Exit(1)
	}

	return theOperation
}

// hex to decimal
func hex(thinp string) (int64, error) {
	return strconv.ParseInt(thinp, 16, 64)
}

// bin to decimal
func bin(thinp string) (int64, error) {
	return strconv.ParseInt(thinp, 2, 64)
}

// bin to decimal
func dec(thinp string) (string, string) {

	n, ero := strconv.ParseInt(thinp, 10, 64)
	if ero != nil {
		fmt.Println(thinp, " is not a valid decimal.")
		takeInput()
	}

	bin := strconv.FormatInt(n, 2)
	hex := strconv.FormatInt(n, 16)

	hex = strings.ToUpper(hex)

	return bin, hex
}
