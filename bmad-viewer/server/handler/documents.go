package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"bmad-viewer/server/model"
)

// DocumentHandler holds the document cache and handles HTTP requests.
type DocumentHandler struct {
	docs []model.Document
}

// NewDocumentHandler creates a handler with the given document cache.
func NewDocumentHandler(docs []model.Document) *DocumentHandler {
	return &DocumentHandler{docs: docs}
}

// HandleList returns all documents as summaries (without content).
func (h *DocumentHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	summaries := make([]model.DocumentSummary, len(h.docs))
	for i := range h.docs {
		summaries[i] = h.docs[i].ToSummary()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summaries)
}

// HandleGet returns a single document by path.
func (h *DocumentHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	docPath := strings.TrimPrefix(r.URL.Path, "/api/documents/")
	if docPath == "" {
		h.HandleList(w, r)
		return
	}

	for i := range h.docs {
		if h.docs[i].Path == docPath {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(h.docs[i])
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "document not found"})
}
