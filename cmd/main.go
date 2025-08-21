package main

import (
	"log"

	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/data"
	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/markdown"
)

func main() {
	commands := data.Commands()

	err := markdown.GenerateFile(commands, "commands.md")
	if err != nil {
		log.Fatalf("cannot generate markdown: %v", err)
	}

	log.Println("Markdown table generated: commands.md")
}
