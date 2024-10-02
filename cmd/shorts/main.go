package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NTUT-NPC/shorts"
	"github.com/NTUT-NPC/shorts/web"
	"github.com/a-h/templ"
	"github.com/fsnotify/fsnotify"
)

func main() {
	shorts.ReadRedirects()
	go watchRedirectsFile()

	shorts.ReadStats()

	http.HandleFunc("/", handleRedirect)
	http.HandleFunc("/api", shorts.EditConfigHandler)
	http.Handle("/web", templ.Handler(web.Index()))

	log.Fatal(http.ListenAndServe(":8080", nil))
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
		if event.Has(fsnotify.Write) && event.Name == shorts.RedirectsFile {
			go shorts.ReadRedirects()
		}
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[1:] // Remove the leading slash

	if url, ok := shorts.Redirects.Permanent[slug]; ok {
		log.Printf("Permanently redirecting %s", slug)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		shorts.UpdateStat(slug)
		return
	}

	if url, ok := shorts.Redirects.Temporary[slug]; ok {
		log.Printf("Temporary redirecting %s", slug)
		http.Redirect(w, r, url, http.StatusFound)
		shorts.UpdateStat(slug)
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
