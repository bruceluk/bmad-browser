package parser

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"bmad-viewer/server/model"
)

// ParseCSV reads a bmad-help.csv file and returns all workflow steps.
func ParseCSV(filePath string) []model.WorkflowStep {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Warning: cannot open CSV %s: %v", filePath, err)
		return nil
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Warning: failed to parse CSV %s: %v", filePath, err)
		return nil
	}

	if len(records) < 2 {
		log.Printf("Warning: CSV %s has no data rows", filePath)
		return nil
	}

	var steps []model.WorkflowStep
	for _, row := range records[1:] { // skip header
		if len(row) < 16 {
			continue
		}

		// Only parse bmm module workflows
		if row[0] != "bmm" {
			continue
		}

		seq, _ := strconv.Atoi(row[4])
		required := strings.ToLower(row[7]) == "true"

		var outputs []string
		if row[15] != "" {
			outputs = strings.Split(row[15], "|")
			for i := range outputs {
				outputs[i] = strings.TrimSpace(outputs[i])
			}
		}

		steps = append(steps, model.WorkflowStep{
			Name:        row[2],
			Code:        row[3],
			Command:     row[6],
			AgentName:   row[10],
			AgentIcon:   row[11],
			Phase:       row[1],
			Required:    required,
			Description: row[13],
			Outputs:     outputs,
			Duration:    estimateDuration(row[2]),
			Sequence:    seq,
		})
	}

	log.Printf("Parsed %d BMM workflow steps from %s", len(steps), filePath)
	return steps
}

// BuildRoleFlows organizes workflow steps into role-based flows.
func BuildRoleFlows(steps []model.WorkflowStep) []model.RoleFlow {
	// Categorize steps by role
	var pmSteps, devSteps, qaSteps []model.WorkflowStep

	for _, s := range steps {
		// Skip "anytime" phase workflows
		if s.Phase == "anytime" {
			continue
		}

		switch {
		case isQAStep(s):
			qaSteps = append(qaSteps, s)
		case isPMStep(s):
			pmSteps = append(pmSteps, s)
		case isDevStep(s):
			devSteps = append(devSteps, s)
		}
	}

	// Sort each by sequence
	sortBySequence(pmSteps)
	sortBySequence(devSteps)
	sortBySequence(qaSteps)

	// Build upstream/downstream relationships
	pmUpstream := []model.WorkflowStep{}
	pmDownstream := firstN(devSteps, 2)

	devUpstream := lastN(pmSteps, 2)
	devDownstream := firstN(qaSteps, 1)

	qaUpstream := lastN(devSteps, 2)
	qaDownstream := []model.WorkflowStep{}

	return []model.RoleFlow{
		{
			Role:       "pm",
			RoleColor:  "#81c784",
			Label:      "产品经理",
			Steps:      pmSteps,
			Upstream:   pmUpstream,
			Downstream: pmDownstream,
		},
		{
			Role:       "developer",
			RoleColor:  "#4fc3f7",
			Label:      "开发者",
			Steps:      devSteps,
			Upstream:   devUpstream,
			Downstream: devDownstream,
		},
		{
			Role:       "qa",
			RoleColor:  "#b39ddb",
			Label:      "测试人员",
			Steps:      qaSteps,
			Upstream:   qaUpstream,
			Downstream: qaDownstream,
		},
	}
}

func isPMStep(s model.WorkflowStep) bool {
	return s.Phase == "1-analysis" || s.Phase == "2-planning"
}

func isDevStep(s model.WorkflowStep) bool {
	if s.Phase != "3-solutioning" && s.Phase != "4-implementation" {
		return false
	}
	// Exclude QA, SM, and retrospective steps
	nameLower := strings.ToLower(s.Name)
	if strings.Contains(nameLower, "qa") || strings.Contains(nameLower, "retrospective") {
		return false
	}
	if strings.Contains(nameLower, "sprint planning") || strings.Contains(nameLower, "sprint status") {
		return false
	}
	if strings.Contains(nameLower, "readiness") {
		return false
	}
	agentLower := strings.ToLower(s.AgentName)
	if strings.Contains(agentLower, "qa") || strings.Contains(agentLower, "quinn") {
		return false
	}
	// Exclude Scrum Master (Bob) steps
	if strings.Contains(agentLower, "bob") {
		return false
	}
	return true
}

func isQAStep(s model.WorkflowStep) bool {
	nameLower := strings.ToLower(s.Name)
	agentLower := strings.ToLower(s.AgentName)
	return strings.Contains(nameLower, "qa") || strings.Contains(agentLower, "qa") || strings.Contains(agentLower, "quinn")
}

func sortBySequence(steps []model.WorkflowStep) {
	sort.Slice(steps, func(i, j int) bool {
		return steps[i].Sequence < steps[j].Sequence
	})
}

func firstN(steps []model.WorkflowStep, n int) []model.WorkflowStep {
	if len(steps) <= n {
		return steps
	}
	return steps[:n]
}

func lastN(steps []model.WorkflowStep, n int) []model.WorkflowStep {
	if len(steps) <= n {
		return steps
	}
	return steps[len(steps)-n:]
}

func estimateDuration(name string) string {
	nameLower := strings.ToLower(name)
	switch {
	case strings.Contains(nameLower, "brainstorm"):
		return "~20min"
	case strings.Contains(nameLower, "brief"):
		return "~15min"
	case strings.Contains(nameLower, "prd"):
		return "~30min"
	case strings.Contains(nameLower, "ux"):
		return "~30min"
	case strings.Contains(nameLower, "architecture"):
		return "~30min"
	case strings.Contains(nameLower, "epic"):
		return "~20min"
	case strings.Contains(nameLower, "sprint"):
		return "~10min"
	case strings.Contains(nameLower, "dev story"):
		return "~45min"
	case strings.Contains(nameLower, "create story"):
		return "~15min"
	case strings.Contains(nameLower, "code review"):
		return "~15min"
	case strings.Contains(nameLower, "qa"):
		return "~20min"
	case strings.Contains(nameLower, "readiness"):
		return "~15min"
	default:
		return "~15min"
	}
}
