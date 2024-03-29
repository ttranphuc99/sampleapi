package app

import (
	"encoding/json"
	"log"
	"net/http"
	"sampleapi/app/models"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format JSON, err=%v \n", err)
	}
}

func mapToJSON(p *models.Post) models.JsonPost {
	return models.JsonPost{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}
