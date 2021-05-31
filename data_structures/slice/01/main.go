package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	data := []string{"hello", "nigga", "rachi"}
	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}

}
