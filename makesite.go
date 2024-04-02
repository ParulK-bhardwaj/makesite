package main

import (
	"fmt"
	"html/template"

	// "io/ioutil" >> Deprecated, replaced with "os"
	"os"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	fileContents, err := os.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	fmt.Print(string(fileContents))

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	// .Must means that any error will result in a panic
	tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	tmpl.Execute(os.Stdout, string(fileContents))

}
