package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
)

type Redirects struct {
	Temporary map[string]string
	Permanent map[string]string
}

var redirects Redirects

const redirectsFile = "config/redirects.toml"

func main() {
	readRedirects()
	go watchRedirectsFile()

	readStats()

	http.HandleFunc("/", handleRedirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
		if event.Has(fsnotify.Write) && event.Name == redirectsFile {
			go readRedirects()
		}
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:] // Remove the leading slash

	if url, ok := redirects.Permanent[slug]; ok {
		log.Printf("Permanently redirecting %s", slug)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		updateStat(slug)
		return
	}

	if url, ok := redirects.Temporary[slug]; ok {
		log.Printf("Temporary redirecting %s", slug)
		http.Redirect(w, r, url, http.StatusFound)
		updateStat(slug)
		return
	}

	// Read and serve custom 404 page, fallback to default if not found
	contents, err := os.ReadFile("config/404.html")
	if err != nil {
		contents, _ = os.ReadFile("assets/404.html")
	}
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(contents)
}
