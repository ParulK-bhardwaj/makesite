package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	Content string
}

func main() {
	// flag.String: This creates a new flag of type string.
	// It takes three arguments:
	// the flag name, the default value, and the usage description.
	dirPtr := flag.String("dir", ".", "The directory to find all .txt files")
	flag.Parse()

	// Initialize a counter for the number of HTML Files
	fileCount := 0

	// Initialize a counter for the total size of HTML Files
	var totalHTMLFileSize int64 = 0

	files, err := ioutil.ReadDir(*dirPtr)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		fmt.Println(file.Name())
		fileContents, err := ioutil.ReadFile(filepath.Join(*dirPtr, file.Name()))
		if err != nil {
			panic(err)
		}

		page := Page{
			Content: string(fileContents),
		}

		// Create a new template in memory named "template.tmpl".
		// When the template is executed, it will parse template.tmpl,
		// looking for {{ }} where we can inject content.
		// .Must means that any error will result in a panic.
		tmpl := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

		baseName := strings.TrimSuffix(file.Name(), ".txt")
		newHtmlFileName := baseName + ".html"

		newFile, err := os.Create(newHtmlFileName)
		if err != nil {
			panic(err)
		}

		// Executing the template injects the Page instance's data,
		// allowing us to render the content of our text file.
		// Furthermore, upon execution, the rendered template will be
		// saved inside the new file we created earlier.
		err = tmpl.Execute(newFile, page)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Generated HTML file: %s\n", newHtmlFileName)

		totalHTMLFileSize += file.Size()
		fileCount++
	}

	// total file size in kilobytes
	// Always return one significant digit after the decimal point.
	totalFileSizeKB := float64(totalHTMLFileSize) / 1024

	// success message : Success! in bold green and count is bold
	fmt.Printf("\033[1;32mSuccess!\033[0m Generated \033[1m%d\033[0m pages (%.1fkB total).\n", fileCount, totalFileSizeKB)
}
