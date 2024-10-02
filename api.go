package shorts

import (
	"encoding/json"
	"net/http"
)

type UpdateRequest struct {
	Slug      string `json:"slug"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Overwrite bool   `json:"overwrite"`
}

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

	switch updateReq.Slug {
	case "temporary":
		_, keyExists := Redirects.Temporary[updateReq.Key]
		if updateReq.Overwrite || !keyExists {
			Redirects.Temporary[updateReq.Key] = updateReq.Value
		}
	case "permanent":
		_, keyExists := Redirects.Permanent[updateReq.Key]
		if updateReq.Overwrite || !keyExists {
			Redirects.Permanent[updateReq.Key] = updateReq.Value
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
