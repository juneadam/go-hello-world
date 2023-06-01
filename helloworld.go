// helloWorld is my first package for learning Go concepts - namely, creating a package name
// to handle all related files in a directory 
// (per Exercism: Go applications are organized in packages. A package is a collection of source files located in the same directory. 
// 	All source files in a directory must share the same package name. It is conventional for the package name to be the 
// 	last directory in the import path. For example, the files in the "math/rand" package begin with the statement package rand.
// 	When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter 
// 	can be used / accessed. The recommended style of naming in Go is that identifiers will be named using camelCase, 
// 	except for those meant to be accessible across packages which should be PascalCase.)
package main
import (
	"fmt"
	"strings"
)


// Main takes in a string, and concatenates "Hello, " with that string, trimming 
// off any whitespace from the beginning or end of the string.
func Hello(inputText string) string {
	var outputText = string("Hello, " + strings.Trim(inputText, " "))
	fmt.Println(outputText)
	return outputText
}
