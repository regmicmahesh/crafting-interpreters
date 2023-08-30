package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {

	os.Mkdir("ast", 0755)
	output_file := "ast/ast_gen.go"
	output, err := os.OpenFile(output_file, os.O_CREATE|os.O_WRONLY, 0644)

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	// parse template
	tmpl, err := template.New("ast").Funcs(funcMap).ParseFiles("templates/ast.go.tmpl")

	if err != nil {
		panic(err)
	}

	types := &map[string][][]string{
		"Literal": {
			{"Value", "any"},
		},
		"Assign": {
			{"Name", "*token.Token"},
			{"Value", "Expr"},
		},
		"Binary": {
			{"Left", "Expr"},
			{"Operator", "*token.Token"},
			{"Right", "Expr"},
		},
	}

	err = tmpl.ExecuteTemplate(output, "ast.go.tmpl", types)

}
