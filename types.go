package activiti

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TaskAction string

const (
	TASK_ACTION_COMPLETE   TaskAction = "complete"
	TASK_ACTION_CLAIM      TaskAction = "claim"
	TASK_ACTION_DELEGATE   TaskAction = "delegate"
	TASK_ACTION_RESOLVE    TaskAction = "resolve"
)

type (
	// JSONTime overrides MarshalJson method to format in ISO8601
	JSONTime time.Time

	// Client represents a Activiti 6.x REST API Client
	ActClient struct {
		Client         *http.Client
		Username       string
		Password       string
		BaseURL        string
		Log            io.Writer // If user set log file name all requests will be logged there
	}

	expirationTime int64

	// ErrorResponse https://www.activiti.org/userguide/#_error_response_body
	ActErrorResponse struct {
		Response        *http.Response        `json:"-"`
		statusCode      string                `json:"statusCode"`
		errorMessage    string                `json:"errorMessage"`
	}

	ActProcessDefinition struct {
		ID                         string   `json:"id,omitempty"`
		URL                        string   `json:"url,omitempty"`
		Version                    int      `json:"version,omitempty"`
		Key                        string   `json:"key,omitempty"`
		Category                   string   `json:"category,omitempty"`
		Suspended                  bool     `json:"suspended,omitempty"`
		Name                       string   `json:"name,omitempty"`
		Description                string   `json:"description,omitempty"`
		DeploymentId               string   `json:"deploymentId,omitempty"`
		DeploymentUrl              string   `json:"deploymentUrl,omitempty"`
		GraphicalNotationDefined   bool     `json:"graphicalNotationDefined,omitempty"`
		Resource                   string   `json:"resource,omitempty"`
		DiagramResource            string   `json:"diagramResource,omitempty"`
		StartFormDefined           bool     `json:"startFormDefined,omitempty"`
	}

	ActProcessDefinitions struct {
		ProcessDefinitions  []ActProcessDefinition    `json:"data,omitempty"`
		Total               int                       `json:"total,omitempty"`
		Start               int                       `json:"start,omitempty"`
		Sort                string                    `json:"sort,omitempty"`
		Order               string                    `json:"order,omitempty"`
		Size                int                       `json:"size,omitempty"`
	}

	ActProcessInstance struct {
		ID                         string   `json:"id,omitempty"`
		URL                        string   `json:"url,omitempty"`
		BusinessKey                string   `json:"businessKey,omitempty"`
		Suspended                  bool     `json:"suspended,omitempty"`
		ProcessDefinitionUrl       string   `json:"processDefinitionUrl,omitempty"`
		ActivityId                 string   `json:"activityId,omitempty"`
		TenantId                   string   `json:"tenantId,omitempty"`
	}

	ActProcessInstances struct {
		ProcessInstances    []ActProcessInstance      `json:"data,omitempty"`
		Total               int                       `json:"total,omitempty"`
		Start               int                       `json:"start,omitempty"`
		Sort                string                    `json:"sort,omitempty"`
		Order               string                    `json:"order,omitempty"`
		Size                int                       `json:"size,omitempty"`
	}

	ActStartProcessInstance struct {
		ProcessDefinitionId    string          `json:"processDefinitionId,omitempty"`
		ProcessDefinitionKey   string          `json:"processDefinitionKey,omitempty"`
		Message                string          `json:"message,omitempty"`
		BusinessKey            string          `json:"businessKey,omitempty"`
		TenantId               string          `json:"tenantId,omitempty"`
		Variables              []ActVariable   `json:"variables,omitempty"`
	}

	ActTask struct {
		Assignee            string   `json:"assignee,omitempty"`
		CreateTime          string   `json:"createTime,omitempty"`
		DelegationState     string   `json:"delegationState,omitempty"`
		Description         string   `json:"description,omitempty"`
		DueDate             string   `json:"dueDate,omitempty"`
		Execution           string   `json:"execution,omitempty"`
		Dd                  string   `json:"id,omitempty"`
		Name                string   `json:"name,omitempty"`
		Owner               string   `json:"owner,omitempty"`
		ParentTask          string   `json:"parentTask,omitempty"`
		Priority            int      `json:"priority,omitempty"`
		ProcessDefinition   string   `json:"processDefinition,omitempty"`
		ProcessInstance     string   `json:"processInstance,omitempty"`
		Suspended           bool     `json:"suspended,omitempty"`
		TaskDefinitionKey   string   `json:"taskDefinitionKey,omitempty"`
		Url                 string   `json:"url,omitempty"`
		TenantId            string   `json:"tenantId,omitempty"`
	}

	ActTasks struct {
		Tasks               []ActTask                 `json:"data,omitempty"`
		Total               int                       `json:"total,omitempty"`
		Start               int                       `json:"start,omitempty"`
		Sort                string                    `json:"sort,omitempty"`
		Order               string                    `json:"order,omitempty"`
		Size                int                       `json:"size,omitempty"`
	}

	ActUser struct {
		ID              string   `json:"id,omitempty"`
		FirstName       string   `json:"firstName,omitempty"`
		LastName        string   `json:"lastName,omitempty"`
		Email           string   `json:"email,omitempty"`
		URL             string   `json:"url,omitempty"`
		PictureURL      string   `json:"pictureUrl,omitempty"`
		Password        string   `json:"password,omitempty"`
	}

	ActUsers struct {
		Users       []ActUser    `json:"data,omitempty"`
		Total       int          `json:"total,omitempty"`
		Start       int          `json:"start,omitempty"`
		Sort        string       `json:"sort,omitempty"`
		Order       string       `json:"order,omitempty"`
		Size        int          `json:"size,omitempty"`
	}

	ActVariable struct {
		Name        string   `json:"name,omitempty"`
		Value       string   `json:"value,omitempty"`
		Operation   string   `json:"operation,omitempty"`
		Type        string   `json:"type,omitempty"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ActErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.errorMessage)
}

// MarshalJSON for JSONTime
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, time.Time(t).UTC().Format(time.RFC3339))
	return []byte(stamp), nil
}

func (e *expirationTime) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}
	i, err := n.Int64()
	if err != nil {
		return err
	}
	*e = expirationTime(i)
	return nil
}
