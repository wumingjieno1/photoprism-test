//go:build ignore
// +build ignore

// This generates countries.go by running "go generate"
package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func gen_migrations(name string) {
	if name == "" {
		return
	}

	dialect := strings.ToLower(name)

	type Migration struct {
		ID         string
		Dialect    string
		Statements []string
	}

	var migrations []Migration

	// Folder in which migration files are stored.
	folder := "./" + dialect

	// Returns directory entries sorted by filename.
	files, _ := os.ReadDir(folder)

	fmt.Printf("generating %s...", dialect)

	strToStmts := func(b []byte) (result []string) {
		stmts := bytes.Split(b, []byte(";\n"))
		result = make([]string, 0, len(stmts))

		for i := range stmts {
			if s := bytes.TrimSpace(stmts[i]); len(s) > 0 {
				if s[len(s)-1] != ';' {
					s = append(s, ';')
				}

				result = append(result, string(s))
			}
		}

		return result
	}

	// Read migrations from files.
	for _, file := range files {
		filePath := filepath.Join(folder, file.Name())

		if file.IsDir() {
			continue
		} else if id := strings.SplitN(filepath.Base(file.Name()), ".", 2)[0]; id == "" {
			fmt.Printf("e")
			// Ignore.
		} else if s, err := os.ReadFile(filePath); err == nil && len(s) > 0 {
			fmt.Printf(".")

			migrations = append(migrations, Migration{ID: id, Dialect: dialect, Statements: strToStmts(s)})
		} else {
			fmt.Printf("f")
			fmt.Println(err.Error())
		}
	}

	fmt.Printf(" found %d migrations\n", len(migrations))

	// Create source file from migrations.
	f, err := os.Create(fmt.Sprintf("dialect_%s.go", dialect))

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Render source template.
	migrationsTemplate.Execute(f, struct {
		Name       string
		Migrations []Migration
	}{
		Name:       name,
		Migrations: migrations,
	})
}

func main() {
	gen_migrations("MySQL")
	gen_migrations("SQLite3")
}

var migrationsTemplate = template.Must(template.New("").Parse(`
package migrate

// Generated code, do not edit.

var Dialect{{ print .Name }} = Migrations{
{{- range .Migrations }}
	{
		ID:        {{ printf "%q" .ID }},
		Dialect:   {{ printf "%q" .Dialect }},
		Statements: []string{ {{ range $index, $s := .Statements}}{{if $index}},{{end}}{{ printf "%q" $s }}{{end}} },
	},	
{{- end }}
}`))
