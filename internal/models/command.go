package models

const (
	Ratatui = "[Ratatui](https://github.com/ratatui/ratatui)"
)

type CommandTransition struct {
	Original OriginalCommand
	Rust     []RustCommand
}

type OriginalCommand struct {
	OriginalCommandName     string // Name of the Unix command or third-party utility, e.g., wget, curl, etc.
	OriginalRepositoryURL   string // URL of the original repository or source code file
	OriginalRepositoryStars int    // Number of stars of the original repository, even if it's just a code file
	OriginalRepositoryLoC   int    // Lines of code in the original repository or source file
}

type RustCommand struct {
	RustCommandName     string // Name of the command rewritten in Rust, e.g., wget-rs, curl-rs, etc.
	RustRepositoryURL   string // URL of the repository with the command rewritten in Rust
	RustRepositoryStars int    // Number of stars of the Rust repository
	RustRepositoryLoC   int    // Lines of code in the Rust repository
	TUI                 string // Information about whether the Rust command has a TUI (Text User Interface) and which library it uses
	SponsorsURL         string // URL for sponsoring the Rust project (if available)
}
