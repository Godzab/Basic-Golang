package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
}

func main() {
	gd := struct {
		Name    string
		Surname string
	}{"Godfrey", "Bafana"}

	err := tpl.ExecuteTemplate(os.Stdout, "define-main.gohtml", gd)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func firstThree(itr []string) []string {
	return itr[:3]
}
