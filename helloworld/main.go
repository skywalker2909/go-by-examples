/*
Go source code is organised into packages which is similar to libraries or modules in othet languages.
A package consists of one or more .go source files in a single directory that define what the package
does.

Each source file begins with a package declaration which states which package the file belongs to followed
by a list of other packages that it imports followed by the declarations in the program.

'package main' is a special package as it defines a standalone executable program instead of a library.
*/
package main

/*
'fmt' is a package that comes with the go standard library that has 100s of other packages. It contains functions
for printing and formatting output and taking input.

'import' tells the go compiler what packages are needed by this source file.
*/
import "fmt"

// 'func main()' is also special as it defines the entry point from where the execution of the program begins.
func main() {
	// Println is a fucntion of the fmt package that prints a line of output
	fmt.Println("Hello World!")
}
