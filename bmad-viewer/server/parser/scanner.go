package parser

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"bmad-viewer/server/model"
)

// ScanDocuments recursively scans a directory for Markdown files and returns parsed documents.
func ScanDocuments(rootDir string) []model.Document {
	var docs []model.Document

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Warning: cannot access %s: %v", path, err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			return nil
		}

		doc, err := ParseMarkdown(path)
		if err != nil {
			log.Printf("Warning: failed to parse %s: %v", path, err)
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			relPath = path
		}
		relPath = filepath.ToSlash(relPath)

		doc.Path = relPath
		doc.Phase = inferPhase(relPath)

		docs = append(docs, *doc)
		return nil
	})

	if err != nil {
		log.Printf("Error scanning directory %s: %v", rootDir, err)
	}

	log.Printf("Scanned %d documents from %s", len(docs), rootDir)
	return docs
}

// inferPhase determines the BMAD phase from the file's relative path.
func inferPhase(relPath string) string {
	parts := strings.Split(relPath, "/")
	if len(parts) == 0 {
		return "other"
	}

	dir := parts[0]
	switch {
	case strings.Contains(dir, "brainstorming"):
		return "analysis"
	case strings.Contains(dir, "planning"):
		return "planning"
	case strings.Contains(dir, "implementation"):
		return "implementation"
	default:
		return "other"
	}
}
