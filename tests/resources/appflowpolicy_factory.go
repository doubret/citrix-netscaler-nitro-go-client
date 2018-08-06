package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"testing"
)

func setup_appflowpolicy(t *testing.T, client *nitro.NitroClient) (*nitro.Appflowpolicy, func()) {
	action, actionTearDown := setup_appflowaction(t, client)

	client.AddAppflowaction(*action)

	resource := nitro.Appflowpolicy{
		Name:        "appflowpolicy",
		Action:      action.Name,
		Comment:     "comment",
		Rule:        "TRUE",
		Undefaction: action.Name,
	}

	return &resource, func() {
		client.DeleteAppflowaction(action.Name)

		if actionTearDown != nil {
			actionTearDown()
		}
	}
}
