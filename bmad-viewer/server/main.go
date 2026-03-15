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
	csvFile := flag.String("csv", "../_bmad/_config/bmad-help.csv", "Path to bmad-help.csv")
	flag.Parse()

	// Scan documents at startup
	docs := parser.ScanDocuments(*dir)

	// Parse workflows and build role flows
	workflows := parser.ParseCSV(*csvFile)
	roles := parser.BuildRoleFlows(workflows)
	log.Printf("Built %d role flows", len(roles))

	// Handlers
	docHandler := handler.NewDocumentHandler(docs)
	wfHandler := handler.NewWorkflowHandler(roles, workflows)

	// API routes
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	http.HandleFunc("/api/documents/", docHandler.HandleGet)
	http.HandleFunc("/api/documents", docHandler.HandleList)
	http.HandleFunc("/api/roles", wfHandler.HandleRoles)
	http.HandleFunc("/api/workflows", wfHandler.HandleWorkflows)

	// Static file handler with SPA fallback (default route)
	http.HandleFunc("/", handler.NewStaticHandler(webDistFS))

	log.Printf("BMAD Viewer server starting on :%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
