package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Spilloverpolicy struct {
	Name    string `json:"name"`
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
	Rule    string `json:"rule,omitempty"`
}

func spilloverpolicy_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type SpilloverpolicyUnset struct {
	Name    string `json:"name"`
	Rule    bool   `json:"rule,omitempty"`
	Action  bool   `json:"action,omitempty"`
	Comment bool   `json:"comment,omitempty"`
}

type update_spilloverpolicy struct {
	Name    string `json:"name"`
	Rule    string `json:"rule,omitempty"`
	Action  string `json:"action,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type rename_spilloverpolicy struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_spilloverpolicy_payload struct {
	Resource Spilloverpolicy `json:"spilloverpolicy"`
}

type rename_spilloverpolicy_payload struct {
	Rename rename_spilloverpolicy `json:"spilloverpolicy"`
}

type unset_spilloverpolicy_payload struct {
	Unset SpilloverpolicyUnset `json:"spilloverpolicy"`
}

type update_spilloverpolicy_payload struct {
	Update update_spilloverpolicy `json:"spilloverpolicy"`
}

type get_spilloverpolicy_result struct {
	Results []Spilloverpolicy `json:"spilloverpolicy"`
}

type count_spilloverpolicy_result struct {
	Results []Count `json:"spilloverpolicy"`
}

func (c *NitroClient) AddSpilloverpolicy(resource Spilloverpolicy) error {
	payload := add_spilloverpolicy_payload{
		resource,
	}

	return c.post("spilloverpolicy", "", nil, payload)
}

func (c *NitroClient) RenameSpilloverpolicy(name string, newName string) error {
	payload := rename_spilloverpolicy_payload{
		rename_spilloverpolicy{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("spilloverpolicy", "", qs, payload)
}

func (c *NitroClient) CountSpilloverpolicy() (int, error) {
	var results count_spilloverpolicy_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("spilloverpolicy", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsSpilloverpolicy(key string) (bool, error) {
	var results count_spilloverpolicy_result

	id, qs := spilloverpolicy_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("spilloverpolicy", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListSpilloverpolicy() ([]Spilloverpolicy, error) {
	results := get_spilloverpolicy_result{}

	if err := c.get("spilloverpolicy", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetSpilloverpolicy(key string) (*Spilloverpolicy, error) {
	var results get_spilloverpolicy_result

	id, qs := spilloverpolicy_key_to_id_args(key)

	if err := c.get("spilloverpolicy", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one spilloverpolicy element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("spilloverpolicy element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteSpilloverpolicy(key string) error {
	id, qs := spilloverpolicy_key_to_id_args(key)

	return c.delete("spilloverpolicy", id, qs)
}

func (c *NitroClient) UnsetSpilloverpolicy(unset SpilloverpolicyUnset) error {
	payload := unset_spilloverpolicy_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("spilloverpolicy", "", qs, payload)
}

func (c *NitroClient) UpdateSpilloverpolicy(resource Spilloverpolicy) error {
	payload := update_spilloverpolicy_payload{
		update_spilloverpolicy{
			resource.Name,
			resource.Rule,
			resource.Action,
			resource.Comment,
		},
	}

	return c.put("spilloverpolicy", "", nil, payload)
}
