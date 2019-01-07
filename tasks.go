package activiti

import (
	"errors"
	"fmt"
)

// GetTask retrieves task by ID
// Endpoint: GET runtime/tasks/{taskId}
func (c *ActClient) GetTask(tid string) (*ActTask, error) {
	tk := &ActTask{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/runtime/tasks/", tid), nil)
	if err != nil {
		return tk, err
	}

	if err = c.SendWithBasicAuth(req, tk); err != nil {
		return tk, err
	}

	return tk, nil
}

// GetTasks retrieves all tasks
// Endpoint: GET runtime/tasks
func (c *ActClient) GetTasks() (*ActTasks, error) {
	tks := &ActTasks{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "/runtime/tasks"), nil)
	if err != nil {
		return tks, err
	}

	if err = c.SendWithBasicAuth(req, tks); err != nil {
		return tks, err
	}

	return tks, nil
}

// TaskAction complete/claim/delegate/resolve a task in activiti
// Endpoint: POST runtime/tasks/{taskId}
func (c *ActClient) TaskAction(tid string, action TaskAction, assignee string) error {
	if tid == "" || action == "" {
		return errors.New("Task id and action are required for task action ")
	}
	var params map[string]string
	if TASK_ACTION_CLAIM == action || TASK_ACTION_DELEGATE == action {
		params = map[string]string{
			"action": string(action),
			"assignee": assignee,
		}
	} else {
		params = map[string]string{
			"action": string(action),
		}
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s%s", c.BaseURL, "/runtime/tasks/", tid), params)
	if err != nil {
		return err
	}

	if err = c.SendWithBasicAuth(req, nil); err != nil {
		return err
	}

	return nil
}