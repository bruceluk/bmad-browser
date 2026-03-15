package model

// Document represents a parsed Markdown document with frontmatter metadata.
type Document struct {
	Path        string                 `json:"path"`
	Title       string                 `json:"title"`
	Content     string                 `json:"content"`
	Frontmatter map[string]any `json:"frontmatter"`
	Phase       string                 `json:"phase"`
}

// DocumentSummary is a lightweight version of Document without content, used for list APIs.
type DocumentSummary struct {
	Path  string `json:"path"`
	Title string `json:"title"`
	Phase string `json:"phase"`
}

// ToSummary converts a Document to a DocumentSummary.
func (d *Document) ToSummary() DocumentSummary {
	return DocumentSummary{
		Path:  d.Path,
		Title: d.Title,
		Phase: d.Phase,
	}
}
