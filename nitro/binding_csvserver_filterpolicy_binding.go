package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type CsvserverFilterpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,string,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverFilterpolicyBindingKey struct {
	Name       string
	Policyname string
}

func (resource CsvserverFilterpolicyBinding) ToKey() CsvserverFilterpolicyBindingKey {
	key := CsvserverFilterpolicyBindingKey{
		resource.Name,
		resource.Policyname,
	}

	return key
}

func (key CsvserverFilterpolicyBindingKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = key.Name
	args = append(args, "policyname:"+key.Policyname)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key CsvserverFilterpolicyBindingKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key CsvserverFilterpolicyBindingKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_csvserver_filterpolicy_binding_payload struct {
	Resource CsvserverFilterpolicyBinding `json:"csvserver_filterpolicy_binding"`
}

func (c *NitroClient) AddCsvserverFilterpolicyBinding(resource CsvserverFilterpolicyBinding) error {
	payload := add_csvserver_filterpolicy_binding_payload{
		resource,
	}

	return c.put("csvserver_filterpolicy_binding", "", nil, payload)
}

//      LIST

type list_csvserver_filterpolicy_binding_result struct {
	Results []CsvserverFilterpolicyBinding `json:"csvserver_filterpolicy_binding"`
}

func (c *NitroClient) ListCsvserverFilterpolicyBinding() ([]CsvserverFilterpolicyBinding, error) {
	results := list_csvserver_filterpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
	}

	if err := c.get("csvserver_filterpolicy_binding", "", qs, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      COUNT

type count_csvserver_filterpolicy_binding_result struct {
	Results []Count `json:"csvserver_filterpolicy_binding"`
}

func (c *NitroClient) CountCsvserverFilterpolicyBinding() (int, error) {
	results := count_csvserver_filterpolicy_binding_result{}

	qs := map[string]string{
		"bulkbindings": "yes",
		"count":        "yes",
	}

	if err := c.get("csvserver_filterpolicy_binding", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      READ

type get_csvserver_filterpolicy_binding_result struct {
	Results []CsvserverFilterpolicyBinding `json:"csvserver_filterpolicy_binding"`
}

func (c *NitroClient) GetCsvserverFilterpolicyBinding(key CsvserverFilterpolicyBindingKey) (*CsvserverFilterpolicyBinding, error) {
	var results get_csvserver_filterpolicy_binding_result

	id, qs := key.to_id_filter()

	if err := c.get("csvserver_filterpolicy_binding", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one csvserver_filterpolicy_binding element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("csvserver_filterpolicy_binding element not found")
		}

		return &results.Results[0], nil
	}
}

//      EXISTS

func (c *NitroClient) ExistsCsvserverFilterpolicyBinding(key CsvserverFilterpolicyBindingKey) (bool, error) {
	var results count_csvserver_filterpolicy_binding_result

	id, qs := key.to_id_filter()

	qs["count"] = "yes"

	if err := c.get("csvserver_filterpolicy_binding", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteCsvserverFilterpolicyBinding(key CsvserverFilterpolicyBindingKey) error {
	id, qs := key.to_id_args()

	return c.delete("csvserver_filterpolicy_binding", id, qs)
}
