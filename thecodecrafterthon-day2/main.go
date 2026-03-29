package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Println("dd")
	theMainFunc()
}

func theMainFunc() {

	/* fmt.Println("Enter First Numbe or Quite")
	var first int
	fmt.Scan(&first) */
	fir := firstNum()
	theLeter, err1 := strconv.Atoi(fir)
	if err1 != nil {
		fmt.Println("MUST BE A NUMBER")
		firstNum()
	}

	sec := secon()
	theSec, err2 := strconv.Atoi(sec)
	if err2 != nil {
		fmt.Println("MUST BE A NUMBER")
		secon()
	}

	ope := opera()

	//do opreation
	switch ope {
	case "a":
		fmt.Println("Answer is:", add(theLeter, theSec))
		theMainFunc()
	case "b":
		fmt.Println("Answer is:", sub(theLeter, theSec))
		theMainFunc()
	case "c":
		fmt.Println("Answer is:", Mult(theLeter, theSec))
		theMainFunc()
	case "d":
		if theSec == 0 {
			fmt.Println("CANNOT DIVIDE BY 0!, TRY AGAIN")
			theMainFunc()
		} else {
			fmt.Println("Answer is:", div(theLeter, theSec))
			theMainFunc()
		}
	default:
		opera()
	}
}

// firsrNum
func firstNum() string {

	fmt.Println("Enter First Numbe or h for HELP and q to Quite")

	var first string
	fmt.Scan(&first)

	switch first {
	case "q":
		os.Exit(1)
	case "h":
		help()
		theMainFunc()
	}

	return first
}

// secnumber
func secon() string {
	fmt.Println("Enter Second Numbe or h for HELP and q to Quite")
	var second string
	fmt.Scan(&second)

	switch second {
	case "q":
		os.Exit(1)
	case "h":
		help()
		theMainFunc()
	}

	return second
}

func opera() string {
	fmt.Println("Chose an operation")
	fmt.Println("a. for addition")
	fmt.Println("b.  for subtration")
	fmt.Println("c. for multiplication")
	fmt.Println("d. Division")

	var oper string
	fmt.Scan(&oper)

	return oper
}

// Sub
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// Mult
func Mult(a int, b int) int {
	return a * b
}

// div
func div(a int, b int) float64 {
	return float64(a) / float64(b)
}

func help() {
	fmt.Println("HELP")
	fmt.Println("Enter first number")
	fmt.Println("Enter second number")
	fmt.Println("Chose an operation to get the result")
}
