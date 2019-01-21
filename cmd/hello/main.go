package main // packageName

import "fmt" // fmt package have functions which help to print into OS stuffs

// main function where everything happen
func main() {
	msg := helloWho("World") // Call function hellWho and return to a new var called msg
	fmt.Println(msg)         // print return of var msg and a new line
}

// helloWho is a function which concat its arg who to string Hello and return its concat string
func helloWho(who string) string {
	return "Hello " + who // concat Literal String with string var who and return its result
}
