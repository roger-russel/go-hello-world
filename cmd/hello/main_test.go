package main // Usually same package name of file that test was writen for

import (
	"testing" // Testing must be imported and used into every test
)

//ExampleMain test func main
//See more https://blog.golang.org/examples
func ExampleMain() {
	main() // The sintaxe bellow is to easily get the output and compare
	// Output:
	// Hello World
}

// Test_helloWho is a auto genereated function on Visual Studio
// Pretty cool, isn't?
func Test_helloWho(t *testing.T) {
	type args struct { // args structs that helloWho receive
		who string // arg with the very same name and type that helloWho expect
	}
	tests := []struct { // Test struct which will be iterated for every test
		name string // The Tast Name for differentiate each test
		args args   // args truct that will be passed to each test
		want string // What is wanted as result from test
	}{ // I wrote this part myself
		{
			"First Test Hello World", // Test name
			args{"World"},            // Args for the test, it must be an args struct
			"Hello World",            // The result which I'm waiting for
		},
		{
			"Second Test Hello Moon",
			args{"Moon"},
			"Hello Moon",
		},
		{
			"Third Test Hello Everybody",
			args{"Everybody"},
			"Hello Everybody",
		},
	} // This part bellow was autogenerated too...
	for _, tt := range tests { // iterated over tests struct
		t.Run(tt.name, func(t *testing.T) { // Call helloWho with each args pass into test Struct
			if got := helloWho(tt.args.who); got != tt.want { // Compare if it's results is what is wanted
				t.Errorf("helloWho() = %v, want %v", got, tt.want) // Print fail Message
			}
		})
	}
}
