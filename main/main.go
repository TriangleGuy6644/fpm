package main

import (
	"os"

	"github.com/fatih/color"
)

// FPM: The Font Package Manager :3
func main() {
	args := os.Args
	if len(args) < 2 {
		color.HiBlack("usage: fpm [install|remove|search|list] [font]")
	}

	cmd := args[1]

	switch cmd {
	case "install":
		if len(args) < 3 {
			color.HiBlack("usage: fpm install [font] [--source google|debian]")
			return
		}
		font := args[2]
		source := "google"
		if len(args) > 4 && args[3] == "--source" {
			source = args[4]
		}
		InstallFont(font, source)
	case "remove":
		if len(args) < 3 {
			color.HiBlack("usage: fpm remove [font]")
			return
		}
		font := args[2]
		RemoveFont(font)
	case "list":
		ListFonts()
	case "search":
		if len(args) < 3 {
			color.HiBlack("usage: fpm search [font]")
			return
		}
		term := args[2]
		SearchFont(term)
	default:
		color.HiBlack("unknown command: ", cmd)
	}
}
