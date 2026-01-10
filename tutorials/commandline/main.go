package main

import (
	"fmt"
	"os"
)

func main() {
	var sentence, separator string

	/*
		Command line arguments are available to a program via the os.Args variable.

		It is a slice of strings where each string corresponds to a command line argument
		entered when running this go program.

		The first element in the slice is the name of the go program itself whereas the
		other elements after it are arguments passed to the program during its execution.
	*/
	for i := 1; i < len(os.Args); i++ {
		sentence += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(sentence)
}
