// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"errors"
	"os"
	"text/template"

	flag "github.com/spf13/pflag"
)

func OpenOutput(path string) (*os.File, error) {
	var (
		err    error
		output *os.File
	)

	output, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func main() {
	var (
		tmplFile = flag.StringP("template", "t", "", "code template")
		xmlFile  = flag.StringP("input", "i", "", "XML DBus IDL file")
		destFile  = flag.StringP("output", "o", "", "output file for code generated")
		pkgName  = flag.StringP("package", "p", "", "go package name")
		intfBase = flag.StringP("intf-base", "b", "", "base name of the DBus interface")

		idl    *Idl
		err    error
		output *os.File
	)

	flag.Parse()

	if *tmplFile == "" {
		panic(errors.New("no template file provided"))
	}

	idl, err = NewIdl(*xmlFile)
	if err != nil {
		panic(err)
	}

	output, err = OpenOutput(*destFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	t := NewTemplateState(idl.Node, *pkgName, idl.Name, *intfBase)

	funcMap := template.FuncMap{
		"deduct":     Deduct,
		"increment":  Increment,
		"initial":    Initial,
		"camelcase":  CamelCase,
		"underscore": Underscore,
		"hyphen":     Hyphen,
	}

	tmpl := template.Must(template.New(*tmplFile).Funcs(funcMap).ParseFiles(*tmplFile))

	err = tmpl.Execute(output, t)
	if err != nil {
		panic(err)
	}
}
