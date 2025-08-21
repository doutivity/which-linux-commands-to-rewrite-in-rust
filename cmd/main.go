package main

import (
	"log"
	"os"
	"text/template"

	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/data"
	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/utils"
)

func main() {
	commands := data.Commands()

	funcMap := template.FuncMap{
		"GoogleSearchURL": utils.GoogleSearchURL,
		"GitHubSearchURL": utils.GitHubSearchURL,
	}

	tmpl, err := template.New("table.html").Funcs(funcMap).ParseFiles("templates/table.html")
	if err != nil {
		log.Fatalf("cannot parse template: %v", err)
	}

	outFile, err := os.Create("commands.html")
	if err != nil {
		log.Fatalf("cannot create HTML file: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, commands); err != nil {
		log.Fatalf("cannot execute template: %v", err)
	}

	log.Println("HTML table generated: commands.html")
}
