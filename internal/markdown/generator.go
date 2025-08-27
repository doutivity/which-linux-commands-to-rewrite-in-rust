package markdown

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/models"
	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/utils"
)

func GenerateFile(commands []models.CommandTransition, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "| Command (C) | Stars (C) | LoC (C) | Command (Rust) | Stars (Rust) | Alternatives | Search more |")
	fmt.Fprintln(file, "|-------------|-----------|---------|----------------|--------------|--------------|-------------|")

	for _, cmd := range commands {
		var (
			rustCopy     = slices.Clone(cmd.Rust)
			mainRust     models.RustCommand
			alternatives []models.RustCommand
		)

		sort.Slice(rustCopy, func(i, j int) bool {
			return rustCopy[i].RustRepositoryStars > rustCopy[j].RustRepositoryStars
		})

		if len(rustCopy) > 0 {
			mainRust, alternatives = rustCopy[0], rustCopy[1:]
		}

		var (
			alternativesMarkdown = make([]string, len(alternatives))
		)

		for i, alternative := range alternatives {
			alternativesMarkdown[i] = fmt.Sprintf("[%s](%s)", alternative.RustCommandName, alternative.RustRepositoryURL)
		}

		if len(rustCopy) == 0 {
			// Немає Rust-реалізацій
			fmt.Fprintf(file, "| [%s](%s) | %d | %d |  |  |  | [GitHub](%s) / [Google](%s) |\n",
				cmd.Original.OriginalCommandName,
				cmd.Original.OriginalRepositoryURL,
				cmd.Original.OriginalRepositoryStars,
				cmd.Original.OriginalRepositoryLoC,
				utils.GitHubSearchURL(cmd.Original.OriginalCommandName),
				utils.GoogleSearchURL(cmd.Original.OriginalCommandName),
			)
		} else {
			fmt.Fprintf(file, "| [%s](%s) | %d | %d | [%s](%s) | %d | %s | [GitHub](%s) / [Google](%s) |\n",
				cmd.Original.OriginalCommandName,
				cmd.Original.OriginalRepositoryURL,
				cmd.Original.OriginalRepositoryStars,
				cmd.Original.OriginalRepositoryLoC,
				mainRust.RustCommandName,
				mainRust.RustRepositoryURL,
				mainRust.RustRepositoryStars,
				strings.Join(alternativesMarkdown, ", "),
				utils.GitHubSearchURL(cmd.Original.OriginalCommandName),
				utils.GoogleSearchURL(cmd.Original.OriginalCommandName),
			)
		}
	}

	return nil
}
