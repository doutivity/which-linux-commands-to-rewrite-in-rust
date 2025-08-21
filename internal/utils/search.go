package utils

import "fmt"

func GoogleSearchURL(query string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s+rust+github", query)
}

func GitHubSearchURL(query string) string {
	return fmt.Sprintf("https://github.com/search?q=%s+rust&type=repositories", query)
}
