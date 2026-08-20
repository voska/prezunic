package main

import (
	"github.com/voska/prezunic"
	"github.com/voska/vtexkit/cli"
)

// version is injected at build time via -ldflags.
var version = "dev"

func main() {
	cli.Main(cli.App{
		Store:       prezunic.Store,
		Version:     version,
		Description: "Prezunic supermarket CLI for humans and AI agents.",
	})
}
