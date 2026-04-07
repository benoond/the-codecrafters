# Go Fundamentals

Go is a programming language designed by Google. It is simple with high performance ment to support backend systems, cloud services, and developer tools.


##  Installation
Before you get started download Go from the official website and install it on operating system. 

### Verify installation:
 Type "go version" on the terminal

## Hello World Example
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
## Running Go Program
Run using:

"go run main.go" on the termianl

## Basic Syntax
- Packages
Every Go program starts with a package declaration:
example package main

- Imports
Used to include standard or external libraries:
example import "fmt"


- Functions
Functions are declared using the func keyword:
example:

func add(a int, b int) int {
    return a + b
}

## Variables and Data Types
Variable Declaration
var name string = "John" //Explicit declaration
age := 25 // shorthand


## Basic Data Types
- int, float64
- string
- bool
- Constants
- const Pi = 3.14

## Control Structures
### if, else if and else
example
If-Else
if age > 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

### Switch
example
Switch
switch day {
case "Monday":
    fmt.Println("Start of the week")
default:
    fmt.Println("Other day")
}

## Loop
Only for exists in Go
example
for i := 0; i < 5; i++ {
    fmt.Println(i)
}


## Arrays, Slices, and Maps
### Arrays
example
var arr [3]int = [3]int{1, 2, 3}

### Slices (Dynamic Arrays)
example
slice := []int{1, 2, 3}
slice = append(slice, 4)


### Maps (Key-Value Pairs)
example
m := make(map[string]int)
m["age"] = 30

### Functions
example
- Multiple Return Values
func divide(a, b int) (int, int) {
    return a / b, a % b
}
- Named Returns
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

## Structs and Methods
### Structs
example
type Person struct {
    Name string
    Age  int
}
### Methods
func (p Person) greet() string {
    return "Hello " + p.Name
}

