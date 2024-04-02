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
	Content string
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

	page := Page{
		Content: string(fileContents),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	// .Must means that any error will result in a panic.
	tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Render the contents of `first-post.txt` using Go Templates and print it to stdout.
	err = tmpl.Execute(os.Stdout, string(fileContents))
	if err != nil {
		panic(err)
	}
	// Create a new, blank HTML file.
	newFile, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	tmpl.Execute(newFile, page)

}
