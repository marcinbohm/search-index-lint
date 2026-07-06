package main

import (
	"os"

	"github.com/marcinbohm/search-index-lint/internal/cli"
)

func main() {
	os.Exit(cli.Execute(os.Args[1:], os.Stdout, os.Stderr))
}
