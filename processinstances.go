package activiti

import (
	"errors"
	"fmt"
)

// GetProcessInstance retrieves process instance by ID
// Endpoint: GET runtime/process-instances/{processInstanceId}
func (c *ActClient) GetProcessInstance(pid string) (*ActProcessInstance, error) {
	pi := &ActProcessInstance{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/runtime/process-instances/", pid), nil)
	if err != nil {
		return pi, err
	}

	if err = c.SendWithBasicAuth(req, pi); err != nil {
		return pi, err
	}

	return pi, nil
}

// GetProcessInstances retrieves all process instances
// Endpoint: GET runtime/process-instances
func (c *ActClient) GetProcessInstances() (*ActProcessInstances, error) {
	pis := &ActProcessInstances{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "/runtime/process-instances"), nil)
	if err != nil {
		return pis, err
	}

	if err = c.SendWithBasicAuth(req, pis); err != nil {
		return pis, err
	}

	return pis, nil
}

// startProcessInstance start a process instance in activiti
// Endpoint: POST runtime/process-instances
func (c *ActClient) startProcessInstance(s ActStartProcessInstance) (*ActProcessInstance, error) {
	pi := &ActProcessInstance{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseURL, "/runtime/process-instances"), s)
	if err != nil {
		return pi, err
	}

	if err = c.SendWithBasicAuth(req, pi); err != nil {
		return pi, err
	}

	return pi, nil
}

// Start a process instance by process definition id
func (c *ActClient) StartProcessInstanceById(pid string) (*ActProcessInstance, error) {
	if pid == "" {
		return nil, errors.New("Process definition id is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionId:pid})
}

// Start a process instance by process definition key
func (c *ActClient) StartProcessInstanceByKey(key string) (*ActProcessInstance, error) {
	if key == "" {
		return nil, errors.New("Process definition key is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{ProcessDefinitionKey:key})
}

// Start a process instance by process definition message
func (c *ActClient) StartProcessInstanceByMessage(msg string) (*ActProcessInstance, error) {
	if msg == "" {
		return nil, errors.New("Message is required to start a process instance ")
	}

	return c.startProcessInstance(ActStartProcessInstance{Message:msg})
}