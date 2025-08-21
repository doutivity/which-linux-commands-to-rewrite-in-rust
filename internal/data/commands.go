package data

import (
	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/models"
)

func Commands() []models.CommandTransition {
	return []models.CommandTransition{
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "cat",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/cat.c",
				OriginalRepositoryStars: 4700,
				OriginalRepositoryLoC:   695,
			},
			Rust: []models.RustCommand{
				{
					RustCommandName:     "bat",
					RustRepositoryURL:   "https://github.com/sharkdp/bat",
					RustRepositoryStars: 54048,
					RustRepositoryLoC:   0,
				},
			},
		},
	}
}
