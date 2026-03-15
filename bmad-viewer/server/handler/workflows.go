package handler

import (
	"encoding/json"
	"net/http"

	"bmad-viewer/server/model"
)

// WorkflowHandler holds workflow data and handles HTTP requests.
type WorkflowHandler struct {
	roles     []model.RoleFlow
	workflows []model.WorkflowStep
}

// NewWorkflowHandler creates a handler with role flows and raw workflow data.
func NewWorkflowHandler(roles []model.RoleFlow, workflows []model.WorkflowStep) *WorkflowHandler {
	return &WorkflowHandler{roles: roles, workflows: workflows}
}

// HandleRoles returns all role flows.
func (h *WorkflowHandler) HandleRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.roles)
}

// HandleWorkflows returns all workflow steps.
func (h *WorkflowHandler) HandleWorkflows(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.workflows)
}
