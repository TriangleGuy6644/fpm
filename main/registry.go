package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GoogleFont struct {
	Family   string            `json:"family"`
	Files    map[string]string `json:"files"`
	Category string            `json:"category"`
}

func FetchGoogleFonts() ([]GoogleFont, error) {
	resp, err := http.Get("https://fonts.google.com/metadata/fonts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	data = data[5:] // strip junk prefix

	var parsed struct {
		FamilyMetadataList []GoogleFont `json:"familyMetadataList"`
	}

	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed.FamilyMetadataList, nil
}

func SearchGoogleFonts(term string) {
	fonts, err := FetchGoogleFonts()
	if err != nil {
		fmt.Println("error fetching fonts:", err)
		return
	}

	term = strings.ToLower(term)
	for _, f := range fonts {
		if strings.Contains(strings.ToLower(f.Family), term) {
			fmt.Println(f.Family, "-", f.Category)
		}
	}
}
