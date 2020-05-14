package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "../models"
	storage "../storages"
)

type taskHandler struct {
	store storage.Store
}

func newTaskHandler(store storage.Store) taskHandler {
	return taskHandler{
		store: store,
	}
}

func (h taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCreateTask(w, r)
	case http.MethodDelete:
		h.handleDeleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h taskHandler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	jsDecoder := json.NewDecoder(r.Body)
	err := jsDecoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	respTask, err := h.store.CreateTask(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to create task in store: %v", err)
		return
	}

	writeJSON(w, &respTask)
}

func (h taskHandler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := shiftPath(r.URL.Path)
	if id == "" {
		http.NotFound(w, r)
		return
	}

	err := h.store.DeleteTask(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to delete task from store: %v", err)
		return
	}
}
