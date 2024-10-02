package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Redirects struct {
	Temporary map[string]string
	Permanent map[string]string
}

func readRedirects() {
	file, err := os.ReadFile(redirectsFile)
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal([]byte(file), &redirects)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d temporary and %d permanent redirects", len(redirects.Temporary), len(redirects.Permanent))
}
