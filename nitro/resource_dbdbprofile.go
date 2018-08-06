package nitro

import (
	"fmt"
	"strconv"
	"strings"
)

type Dbdbprofile struct {
	Name                   string `json:"name"`
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
}

func dbdbprofile_key_to_id_args(key string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strings.Join

	qs := map[string]string{}

	return key, qs
}

type DbdbprofileUnset struct {
	Name                   string `json:"name"`
	Interpretquery         bool   `json:"interpretquery,omitempty"`
	Stickiness             bool   `json:"stickiness,omitempty"`
	Kcdaccount             bool   `json:"kcdaccount,omitempty"`
	Conmultiplex           bool   `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff bool   `json:"enablecachingconmuxoff,omitempty"`
}

type update_dbdbprofile struct {
	Name                   string `json:"name"`
	Interpretquery         string `json:"interpretquery,omitempty"`
	Stickiness             string `json:"stickiness,omitempty"`
	Kcdaccount             string `json:"kcdaccount,omitempty"`
	Conmultiplex           string `json:"conmultiplex,omitempty"`
	Enablecachingconmuxoff string `json:"enablecachingconmuxoff,omitempty"`
}

type rename_dbdbprofile struct {
	Name    string `json:"name"`
	Newname string `json:"newname"`
}

type add_dbdbprofile_payload struct {
	Resource Dbdbprofile `json:"dbdbprofile"`
}

type rename_dbdbprofile_payload struct {
	Rename rename_dbdbprofile `json:"dbdbprofile"`
}

type unset_dbdbprofile_payload struct {
	Unset DbdbprofileUnset `json:"dbdbprofile"`
}

type update_dbdbprofile_payload struct {
	Update update_dbdbprofile `json:"dbdbprofile"`
}

type get_dbdbprofile_result struct {
	Results []Dbdbprofile `json:"dbdbprofile"`
}

type count_dbdbprofile_result struct {
	Results []Count `json:"dbdbprofile"`
}

func (c *NitroClient) AddDbdbprofile(resource Dbdbprofile) error {
	payload := add_dbdbprofile_payload{
		resource,
	}

	return c.post("dbdbprofile", "", nil, payload)
}

func (c *NitroClient) RenameDbdbprofile(name string, newName string) error {
	payload := rename_dbdbprofile_payload{
		rename_dbdbprofile{
			name,
			newName,
		},
	}

	qs := map[string]string{
		"action": "rename",
	}

	return c.post("dbdbprofile", "", qs, payload)
}

func (c *NitroClient) CountDbdbprofile() (int, error) {
	var results count_dbdbprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("dbdbprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

func (c *NitroClient) ExistsDbdbprofile(key string) (bool, error) {
	var results count_dbdbprofile_result

	id, qs := dbdbprofile_key_to_id_args(key)

	qs["count"] = "yes"

	if err := c.get("dbdbprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

func (c *NitroClient) ListDbdbprofile() ([]Dbdbprofile, error) {
	results := get_dbdbprofile_result{}

	if err := c.get("dbdbprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

func (c *NitroClient) GetDbdbprofile(key string) (*Dbdbprofile, error) {
	var results get_dbdbprofile_result

	id, qs := dbdbprofile_key_to_id_args(key)

	if err := c.get("dbdbprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one dbdbprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("dbdbprofile element not found")
		}

		return &results.Results[0], nil
	}
}

func (c *NitroClient) DeleteDbdbprofile(key string) error {
	id, qs := dbdbprofile_key_to_id_args(key)

	return c.delete("dbdbprofile", id, qs)
}

func (c *NitroClient) UnsetDbdbprofile(unset DbdbprofileUnset) error {
	payload := unset_dbdbprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.put("dbdbprofile", "", qs, payload)
}

func (c *NitroClient) UpdateDbdbprofile(resource Dbdbprofile) error {
	payload := update_dbdbprofile_payload{
		update_dbdbprofile{
			resource.Name,
			resource.Interpretquery,
			resource.Stickiness,
			resource.Kcdaccount,
			resource.Conmultiplex,
			resource.Enablecachingconmuxoff,
		},
	}

	return c.put("dbdbprofile", "", nil, payload)
}
