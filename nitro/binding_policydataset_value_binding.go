package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type PolicydatasetValueBinding struct {
	Index int    `json:"index,string,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PolicydatasetValueBindingKey struct {
	Name  string
	Value string
}

type add_policydataset_value_binding_payload struct {
	Resources PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

type get_policydataset_value_binding_result struct {
	Results []PolicydatasetValueBinding `json:"policydataset_value_binding"`
}

type count_policydataset_value_binding_result struct {
	Results []Count `json:"policydataset_value_binding"`
}

func policydataset_value_binding_key_to_id_args(key PolicydatasetValueBindingKey) (string, map[string]string) {
	var _ = strconv.Itoa
	var args []string

	args = append(args, "name:"+key.Name)
	args = append(args, "value:"+key.Value)

	qs := map[string]string{}

	if len(args) > 0 {
		qs["args"] = strings.Join(args, ",")
	}

	return "", qs
}

func (c *NitroClient) AddPolicydatasetValueBinding(binding PolicydatasetValueBinding) error {
	payload := add_policydataset_value_binding_payload{
		binding,
	}

	return c.put("policydataset_value_binding", "", nil, payload)
}

func (c *NitroClient) BulkCountPolicydatasetValueBinding() (int, error) {
	var results count_policydataset_value_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("policydataset_value_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) CountPolicydatasetValueBinding(id string) (int, error) {
	var results count_policydataset_value_binding_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("policydataset_value_binding", id, qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsPolicydatasetValueBinding(id string) (bool, error) {
	if count, err := c.CountPolicydatasetValueBinding(id); err != nil {
		return false, err
	} else {
		return count == 1, nil
	}
}

func (c *NitroClient) BulkListPolicydatasetValueBinding() ([]PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding_result

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("policydataset_value_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) ListPolicydatasetValueBinding(id string) ([]PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding_result

	if err := c.get("policydataset_value_binding", id, nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetPolicydatasetValueBinding(key PolicydatasetValueBindingKey) (*PolicydatasetValueBinding, error) {
	var results get_policydataset_value_binding_result

	id, qs := policydataset_value_binding_key_to_id_args(key)

	if err := c.get("policydataset_value_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one policydataset_value_binding element found")
		} else if len(results.Results) < 1 {
			// TODO
			// return nil, fmt.Errorf("policydataset_value_binding element not found")
			return nil, nil
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeletePolicydatasetValueBinding(key PolicydatasetValueBindingKey) error {
	id, qs := policydataset_value_binding_key_to_id_args(key)

	return c.delete("policydataset_value_binding", id, qs)
}
