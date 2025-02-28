package shorts

import (
	"encoding/json"
	"net/http"
)

type UpdateRequest struct {
	Type      string
	Slug      string
	URL       string
	Overwrite bool
}

// EditConfigHandler handles the /api/edit endpoint which allows updating redirect configurations
// It supports modifying both temporary and permanent redirects and can conditionally overwrite existing entries
func EditConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Error: Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateReq UpdateRequest

	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch updateReq.Type {
	case "temporary":
		_, keyExists := Redirects.Temporary[updateReq.Slug]
		if updateReq.Overwrite || !keyExists {
			Redirects.Temporary[updateReq.Slug] = updateReq.URL
		}
	case "permanent":
		_, keyExists := Redirects.Permanent[updateReq.Slug]
		if updateReq.Overwrite || !keyExists {
			Redirects.Permanent[updateReq.Slug] = updateReq.URL
		}
	default:
		http.Error(w, "Invalid section", http.StatusBadRequest)
		return
	}

	err = WriteRedirects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TryRedirectHandler handles the /api/try endpoint which attempts to redirect to a slug
// and falls back to a catch path if the slug is not found
func TryRedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	slug := r.URL.Query().Get("slug")
	if slug == "" {
		http.Error(w, "Missing required parameter: slug", http.StatusBadRequest)
		return
	}

	catch := r.URL.Query().Get("catch")
	if catch == "" {
		http.Error(w, "Missing required parameter: catch", http.StatusBadRequest)
		return
	}

	// Check if slug exists in redirects map
	if url, ok := Redirects.Permanent[slug]; ok {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		UpdateStat(slug)
		return
	}

	if url, ok := Redirects.Temporary[slug]; ok {
		http.Redirect(w, r, url, http.StatusFound)
		UpdateStat(slug)
		return
	}

	// Slug not found, redirect to the fallback URL
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	fallbackURL := scheme + "://" + r.Host + "/" + catch
	http.Redirect(w, r, fallbackURL, http.StatusFound)
}
