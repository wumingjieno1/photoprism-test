//go:build ignore
// +build ignore

// This generates stopwords.go by running "go generate"
package main

import (
	"bufio"
	"io"
	"os"
	"text/template"
)

func main() {
	file, err := os.Open("./resources/stopwords.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	words := []string{}

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		words = append(words, string(line))
	}

	f, err := os.Create("stopwords.go")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	packageTemplate.Execute(f, struct {
		Words []string
	}{
		Words: words,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`
package txt

// Generated code, do not edit.

// StopWords contains a list of stopwords for full-text indexing.
var StopWords = map[string]bool{
{{- range .Words }}
	{{ printf "%q" . }}: true,
{{- end }}
}`))
