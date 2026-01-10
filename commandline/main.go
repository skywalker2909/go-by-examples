package main

// Multiple packages can also be imported using a parenthesized list.
import (
	"fmt"
	// the 'os' package provides functions and other values for dealing with
	// operating systems in a platform independent fashion.
	"os"
)

func main() {
	/*
		'var' is used to declare a variable (of type string in this case).

		if not explicity initialized, then it is implicity initialzied to its
		zero value which is an empty string "" in this case.
	*/
	var sentence, separator string

	/*
		Command line arguments are available to a program via the os.Args variable.

		It is a slice of strings where each string corresponds to a command line argument
		entered when running this go program.

		The first element in the slice is the name of the go program itself whereas the
		other elements after it are arguments passed to the program during its execution.
	*/
	for i := 1; i < len(os.Args); i++ {
		/*
			+ is an arithmetic operator and when applied to strings, it concatenates their
			values to form one big string.

			This is an assignment statment where it concatenates the old value of 'sentence'
			with 'separator + os.Args[i]' and assigns it back to 'sentence'. It can be also
			written as follows:
			sentence = sentence + separator + os.Args[i]ß
		*/
		sentence += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(sentence)
}
