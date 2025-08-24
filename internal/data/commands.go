package data

import (
	"github.com/doutivity/which-linux-commands-to-rewrite-in-rust/internal/models"
)

const (
	coreutilsRepositoryStars = 4700
)

func Commands() []models.CommandTransition {
	return []models.CommandTransition{
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "cat",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/cat.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   695,
			},
			Rust: []models.RustCommand{
				{
					RustCommandName:     "bat",
					RustRepositoryURL:   "https://github.com/sharkdp/bat",
					RustRepositoryStars: 54048,
					RustRepositoryLoC:   0,
					TUI:                 "",
					SponsorsURL:         "https://github.com/sharkdp/bat#sponsors",
				},
			},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "chmod",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/chmod.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   566,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "cp",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/cp.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   1116,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "echo",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/echo.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   247,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "du",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/du.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   963,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "head",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/head.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   930,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "hostname",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/hostname.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   94,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "kill",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/kill.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   269,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "ls",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/ls.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   4821,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "mkdir",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/mkdir.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   267,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "mv",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/mv.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   494,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "pwd",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/pwd.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   333,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "rm",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/rm.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   328,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "sleep",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/sleep.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   122,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "tee",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/tee.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   294,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "tail",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/tail.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   2166,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "touch",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/touch.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   375,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "true",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/true.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   67,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "yes",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/yes.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   113,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "wc",
				OriginalRepositoryURL:   "https://github.com/coreutils/coreutils/blob/master/src/wc.c",
				OriginalRepositoryStars: coreutilsRepositoryStars,
				OriginalRepositoryLoC:   855,
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "ffmpeg",
				OriginalRepositoryURL:   "https://github.com/FFmpeg/FFmpeg",
				OriginalRepositoryStars: 52539,
				OriginalRepositoryLoC:   1538205, // https://codetabs.com/count-loc/count-loc-online.html
			},
			Rust: []models.RustCommand{},
		},
		{
			Original: models.OriginalCommand{
				OriginalCommandName:     "f3",
				OriginalRepositoryURL:   "https://github.com/AltraMayor/f3",
				OriginalRepositoryStars: 2937,
				OriginalRepositoryLoC:   6368, // https://codetabs.com/count-loc/count-loc-online.html
			},
			Rust: []models.RustCommand{},
		},
	}
}
