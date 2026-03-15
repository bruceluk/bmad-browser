package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"bmad-viewer/server/handler"
	"bmad-viewer/server/parser"
)

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	dir := flag.String("dir", "../_bmad-output", "Path to _bmad-output directory")
	flag.Parse()

	// Scan documents at startup
	docs := parser.ScanDocuments(*dir)

	// Document API handler
	docHandler := handler.NewDocumentHandler(docs)

	// API routes (registered first, take priority)
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	http.HandleFunc("/api/documents/", docHandler.HandleGet)
	http.HandleFunc("/api/documents", docHandler.HandleList)

	// Static file handler with SPA fallback (default route)
	http.HandleFunc("/", handler.NewStaticHandler(webDistFS))

	log.Printf("BMAD Viewer server starting on :%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
