package handlers

import (
	"api-postgresql/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao listar tarefas: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
