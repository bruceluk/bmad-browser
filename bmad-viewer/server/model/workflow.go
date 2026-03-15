package model

// WorkflowStep represents a single step in a BMAD workflow.
type WorkflowStep struct {
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	Command     string   `json:"command"`
	AgentName   string   `json:"agentName"`
	AgentIcon   string   `json:"agentIcon"`
	Phase       string   `json:"phase"`
	Required    bool     `json:"required"`
	Description string   `json:"description"`
	Outputs     []string `json:"outputs"`
	Duration    string   `json:"duration"`
	Sequence    int      `json:"sequence"`
}

// RoleFlow represents a role's workflow with core steps and upstream/downstream context.
type RoleFlow struct {
	Role       string         `json:"role"`
	RoleColor  string         `json:"roleColor"`
	Label      string         `json:"label"`
	Steps      []WorkflowStep `json:"steps"`
	Upstream   []WorkflowStep `json:"upstream"`
	Downstream []WorkflowStep `json:"downstream"`
}
