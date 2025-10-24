package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func InstallFont(name string, source string) {
	switch source {
	case "google":
		InstallGoogleFont(name)
	default:
		fmt.Println("unknown source: ", source)
	}
}
func InstallGoogleFont(name string) {
	fonts, err := FetchGoogleFonts()
	if err != nil {
		fmt.Println("error fetching fonts! ", err)
		return
	}

	var fontURL string
	for _, f := range fonts {
		if f.Family == name {
			if u, ok := f.Files["regular"]; ok {
				fontURL = u
			}
			break
		}
	}
	if fontURL == "" {
		fmt.Println("font not found: ", name)
		return
	}

	dir := filepath.Join(os.Getenv("HOME"), ".local/share/fonts", name)
	os.MkdirAll(dir, 0755)
	file := filepath.Join(dir, name+".ttf")
	resp, err := http.Get(fontURL)
	if err != nil {
		fmt.Println("error downloading font: ", err)
		return
	}
	defer resp.Body.Close()

	out, _ := os.Create(file)
	io.Copy(out, resp.Body)
	out.Close()

	color.Green("installed font: ", name, "to ", dir)

}

func RemoveFont(name string) {
	dir := filepath.Join(os.Getenv("HOME"), ".local/share/fonts", name)
	os.RemoveAll(dir)
	color.Red("removed font: ", name)
}

func ListFonts() {
	path := filepath.Join(os.Getenv("HOME"), ".local/share/fonts")
	entries, _ := os.ReadDir(path)
	for _, e := range entries {
		color.Blue(e.Name())
	}
}

func SearchFont(term string) {
	SearchGoogleFonts(term)
}
