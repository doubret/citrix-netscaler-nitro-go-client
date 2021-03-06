package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Netprofile struct {
	Name             string `json:"name,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,string,omitempty"`
}

type NetprofileKey struct {
	Name string
}

func (resource Netprofile) ToKey() NetprofileKey {
	key := NetprofileKey{
		resource.Name,
	}

	return key
}

func (key NetprofileKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Name)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key NetprofileKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key NetprofileKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_netprofile_payload struct {
	Resource Netprofile `json:"netprofile"`
}

func (c *NitroClient) AddNetprofile(resource Netprofile) error {
	payload := add_netprofile_payload{
		resource,
	}

	return c.post("netprofile", "", nil, payload)
}

//      LIST

type list_netprofile_result struct {
	Results []Netprofile `json:"netprofile"`
}

func (c *NitroClient) ListNetprofile() ([]Netprofile, error) {
	results := list_netprofile_result{}

	if err := c.get("netprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_netprofile_result struct {
	Results []Netprofile `json:"netprofile"`
}

func (c *NitroClient) GetNetprofile(key NetprofileKey) (*Netprofile, error) {
	var results get_netprofile_result

	id, qs := key.to_id_args()

	if err := c.get("netprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one netprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("netprofile element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_netprofile_result struct {
	Results []Count `json:"netprofile"`
}

func (c *NitroClient) CountNetprofile() (int, error) {
	var results count_netprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("netprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsNetprofile(key NetprofileKey) (bool, error) {
	var results count_netprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("netprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteNetprofile(key NetprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("netprofile", id, qs)
}

//      UPDATE

type NetprofileUpdate struct {
	Name             string `json:"name,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
}

func (resource Netprofile) ToUpdate() NetprofileUpdate {
	update := NetprofileUpdate{
		resource.Name,
		resource.Srcip,
		resource.Srcippersistency,
		resource.Overridelsn,
	}

	return update
}

type update_netprofile_payload struct {
	Update NetprofileUpdate `json:"netprofile"`
}

func (c *NitroClient) UpdateNetprofile(update NetprofileUpdate) error {
	payload := update_netprofile_payload{
		update,
	}

	return c.put("netprofile", "", nil, payload)
}

//      UNSET

type NetprofileUnset struct {
	Name             string `json:"name,omitempty"`
	Srcip            bool   `json:"srcip,omitempty"`
	Srcippersistency bool   `json:"srcippersistency,omitempty"`
	Overridelsn      bool   `json:"overridelsn,omitempty"`
}

func (resource Netprofile) ToUnset() NetprofileUnset {
	unset := NetprofileUnset{
		resource.Name,
		false,
		false,
		false,
	}

	return unset
}

type unset_netprofile_payload struct {
	Unset NetprofileUnset `json:"netprofile"`
}

func (c *NitroClient) UnsetNetprofile(unset NetprofileUnset) error {
	payload := unset_netprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("netprofile", "", qs, payload)
}

//      RENAME
//      TODO
