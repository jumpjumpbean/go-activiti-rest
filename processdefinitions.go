package activiti

import (
	"fmt"
)

// GetUser retrieves user by ID
// Endpoint: GET repository/process-definitions/{processDefinitionId}
func (c *ActClient) GetProcessDefinition(pid string) (*ActProcessDefinition, error) {
	pd := &ActProcessDefinition{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/repository/process-definitions/", pid), nil)
	if err != nil {
		return pd, err
	}

	if err = c.SendWithBasicAuth(req, pd); err != nil {
		return pd, err
	}

	return pd, nil
}

// GetProcessDefinitions retrieves all process definitions
// Endpoint: GET repository/process-definitions
func (c *ActClient) GetProcessDefinitions() (*ActProcessDefinitions, error) {
	pds := &ActProcessDefinitions{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "repository/process-definitions"), nil)
	if err != nil {
		return pds, err
	}

	if err = c.SendWithBasicAuth(req, pds); err != nil {
		return pds, err
	}

	return pds, nil
}