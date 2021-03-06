package nitro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Lbprofile struct {
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
}

type LbprofileKey struct {
	Lbprofilename string
}

func (resource Lbprofile) ToKey() LbprofileKey {
	key := LbprofileKey{
		resource.Lbprofilename,
	}

	return key
}

func (key LbprofileKey) to_id_params(qsKey string) (string, map[string]string) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	var id string
	var args []string

	id = url.QueryEscape(key.Lbprofilename)

	qs := map[string]string{}

	if len(args) > 0 {
		qs[qsKey] = strings.Join(args, ",")
	}

	return id, qs
}

func (key LbprofileKey) to_id_args() (string, map[string]string) {
	return key.to_id_params("args")
}

func (key LbprofileKey) to_id_filter() (string, map[string]string) {
	return key.to_id_params("filter")
}

//      CREATE

type add_lbprofile_payload struct {
	Resource Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) AddLbprofile(resource Lbprofile) error {
	payload := add_lbprofile_payload{
		resource,
	}

	return c.post("lbprofile", "", nil, payload)
}

//      LIST

type list_lbprofile_result struct {
	Results []Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) ListLbprofile() ([]Lbprofile, error) {
	results := list_lbprofile_result{}

	if err := c.get("lbprofile", "", nil, &results); err != nil {
		return nil, err
	} else {
		return results.Results, err
	}
}

//      READ

type get_lbprofile_result struct {
	Results []Lbprofile `json:"lbprofile"`
}

func (c *NitroClient) GetLbprofile(key LbprofileKey) (*Lbprofile, error) {
	var results get_lbprofile_result

	id, qs := key.to_id_args()

	if err := c.get("lbprofile", id, qs, &results); err != nil {
		return nil, err
	} else {
		if len(results.Results) > 1 {
			return nil, fmt.Errorf("More than one lbprofile element found")
		} else if len(results.Results) < 1 {
			return nil, fmt.Errorf("lbprofile element not found")
		}

		return &results.Results[0], nil
	}
}

//      COUNT

type count_lbprofile_result struct {
	Results []Count `json:"lbprofile"`
}

func (c *NitroClient) CountLbprofile() (int, error) {
	var results count_lbprofile_result

	qs := map[string]string{
		"count": "yes",
	}

	if err := c.get("lbprofile", "", qs, &results); err != nil {
		return -1, err
	} else {
		return results.Results[0].Count, err
	}
}

//      EXISTS

func (c *NitroClient) ExistsLbprofile(key LbprofileKey) (bool, error) {
	var results count_lbprofile_result

	id, qs := key.to_id_args()

	qs["count"] = "yes"

	if err := c.get("lbprofile", id, qs, &results); err != nil {
		// TODO : detect 404
		// return false, err
		return false, nil
	} else {
		return results.Results[0].Count == 1, nil
	}
}

//      DELETE

func (c *NitroClient) DeleteLbprofile(key LbprofileKey) error {
	id, qs := key.to_id_args()

	return c.delete("lbprofile", id, qs)
}

//      UPDATE

type LbprofileUpdate struct {
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
}

func (resource Lbprofile) ToUpdate() LbprofileUpdate {
	update := LbprofileUpdate{
		resource.Lbprofilename,
		resource.Cookiepassphrase,
		resource.Dbslb,
		resource.Processlocal,
		resource.Httponlycookieflag,
		resource.Usesecuredpersistencecookie,
		resource.Useencryptedpersistencecookie,
	}

	return update
}

type update_lbprofile_payload struct {
	Update LbprofileUpdate `json:"lbprofile"`
}

func (c *NitroClient) UpdateLbprofile(update LbprofileUpdate) error {
	payload := update_lbprofile_payload{
		update,
	}

	return c.put("lbprofile", "", nil, payload)
}

//      UNSET

type LbprofileUnset struct {
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Cookiepassphrase              bool   `json:"cookiepassphrase,omitempty"`
	Dbslb                         bool   `json:"dbslb,omitempty"`
	Processlocal                  bool   `json:"processlocal,omitempty"`
	Httponlycookieflag            bool   `json:"httponlycookieflag,omitempty"`
	Usesecuredpersistencecookie   bool   `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie bool   `json:"useencryptedpersistencecookie,omitempty"`
}

func (resource Lbprofile) ToUnset() LbprofileUnset {
	unset := LbprofileUnset{
		resource.Lbprofilename,
		false,
		false,
		false,
		false,
		false,
		false,
	}

	return unset
}

type unset_lbprofile_payload struct {
	Unset LbprofileUnset `json:"lbprofile"`
}

func (c *NitroClient) UnsetLbprofile(unset LbprofileUnset) error {
	payload := unset_lbprofile_payload{
		unset,
	}

	qs := map[string]string{
		"action": "unset",
	}

	return c.post("lbprofile", "", qs, payload)
}

//      RENAME
//      TODO
