package utils

import (
	"fmt"
	"net/url"
)

func GoogleSearchURL(commandName string) string {
	values := url.Values{
		"q": {fmt.Sprintf("site:github.com %q AND %q", commandName, "Rust")},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func GitHubSearchURL(commandName string) string {
	// https://github.com/search?q=%22cat%22+language%3ARust&type=repositories&l=Rust

	values := url.Values{
		"q":    {fmt.Sprintf("%q language:Rust", commandName)},
		"type": {"repositories"},
		"l":    {"Rust"},
	}

	return "https://github.com/search?" + values.Encode()
}
