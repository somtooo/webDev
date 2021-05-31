package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

var tpl = template.New("")

func init() {
	tpl = template.Must(tpl.Funcs(fm).ParseFiles("tpl.html"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
