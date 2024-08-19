package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml/v2"
)

type Redirects struct {
	Temporary map[string]string
	Permanent map[string]string
}

var redirects Redirects

func main() {
	updateRedirects()
	go watchRedirectsFile()

	http.HandleFunc("/", handleRedirect)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func updateRedirects() {
	file, err := os.ReadFile("config/redirects.toml")
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal([]byte(file), &redirects)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d temporary and %d permanent redirects", len(redirects.Temporary), len(redirects.Permanent))
}

func watchRedirectsFile() {
	// Create new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Add config directory to watcher
	err = watcher.Add("config")
	if err != nil {
		log.Fatal(err)
	}

	// Watch for events
	for event := range watcher.Events {
		if event.Has(fsnotify.Write) && event.Name == "config/redirects.toml" {
			go updateRedirects()
		}
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:] // Remove the leading slash

	log.Printf("Redirecting %s", slug)

	if url, ok := redirects.Permanent[slug]; ok {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	}

	if url, ok := redirects.Temporary[slug]; ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	http.NotFound(w, r)
}
