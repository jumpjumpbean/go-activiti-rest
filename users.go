package activiti

import (
	"fmt"
)

// GetUser retrieves user by ID
// Endpoint: GET identity/users/{userId}
func (c *ActClient) GetUser(uid string) (*ActUser, error) {
	user := &ActUser{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.BaseURL, "/identity/users/", uid), nil)
	if err != nil {
		return user, err
	}

	if err = c.SendWithBasicAuth(req, user); err != nil {
		return user, err
	}

	return user, nil
}

// GetUsers retrieves all users
// Endpoint: GET identity/users
func (c *ActClient) GetUsers() (*ActUsers, error) {
	users := &ActUsers{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, "/identity/users"), nil)
	if err != nil {
		return users, err
	}

	if err = c.SendWithBasicAuth(req, users); err != nil {
		return users, err
	}

	return users, nil
}

// CreateUser creates a user in activiti
// Endpoint: POST identity/users
func (c *ActClient) CreateUser(u ActUser) (*ActUser, error) {
	user := &ActUser{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseURL, "/identity/users"), u)
	if err != nil {
		return user, err
	}

	if err = c.SendWithBasicAuth(req, user); err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser updates a user in activiti
// Endpoint: PUT identity/users/{userId}
func (c *ActClient) UpdateUser(u ActUser) (*ActUser, error) {
	user := &ActUser{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s%s", c.BaseURL, "/identity/users/", u.ID), u)
	if err != nil {
		return user, err
	}

	if err = c.SendWithBasicAuth(req, user); err != nil {
		return user, err
	}

	return user, nil
}

// DeleteUser deletes a user in activiti
// Endpoint: DELETE identity/users/{userId}
func (c *ActClient) DeleteUser(uid string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s%s", c.BaseURL, "/identity/users/", uid), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithBasicAuth(req, nil); err != nil {
		return err
	}

	return nil
}