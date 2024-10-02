package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type UpdateRequest struct {
	Section   string `json:"section"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Overwrite bool   `json:"overwrite"`
}

func editConfigHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Error: Bad Request", http.StatusBadRequest)
		return
	}

	switch updateReq.Section {
	case "temporary":
		_, keyExists := redirects.Temporary[updateReq.Key]
		if updateReq.Overwrite || !keyExists {
			redirects.Temporary[updateReq.Key] = updateReq.Value
		}
	case "permanent":
		_, keyExists := redirects.Permanent[updateReq.Key]
		if updateReq.Overwrite || !keyExists {
			redirects.Permanent[updateReq.Key] = updateReq.Value
		}
	default:
		http.Error(w, "Invalid section", http.StatusBadRequest)
		return
	}

	file, err := os.Create(redirectsFile)
	if err != nil {
		http.Error(w, "Error writing config file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(redirects)
	if err != nil {
		http.Error(w, "Error encoding config file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Config updated successfully\n"))
}
