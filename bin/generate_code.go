package main

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"strings"
	"text/template"
)

var types = []string{
	"int",
	"float64",
	"bool",
	"string",
	"AnyValue",
}

var templateFuncs = template.FuncMap{
	"title": strings.Title,
}

func main() {
	executeTemplate(
		os.Stdout,
		templateFrom(os.Stdin),
	)
}

func templateFrom(reader io.Reader) *template.Template {
	input := make([]byte, 10*1024)
	n, err := reader.Read(input)
	if err != nil {
		panic(err)
	}
	input = input[:n]

	output, err := template.New("template").Funcs(templateFuncs).Parse(string(input))
	if err != nil {
		panic(err)
	}
	return output
}

func executeTemplate(writer io.Writer, tmpl *template.Template) {
	// execute template
	var output bytes.Buffer
	err := tmpl.Execute(&output, types)
	if err != nil {
		panic(err)
	}

	// format go code (ie, "go fmt")
	formatted, err := format.Source(output.Bytes())
	if err != nil {
		panic(err)
	}

	// ...and output it
	writer.Write(formatted)
}
