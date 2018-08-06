package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNstcpprofile(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_nstcpprofile(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddNstcpprofile(*resource)

	assert.NoError(t, err)

	log.Print("--EXISTS--")
	exists, err := client.ExistsNstcpprofile(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	log.Print("--EXISTS--")
	exists, err = client.ExistsNstcpprofile(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--GET--")
	res, err := client.GetNstcpprofile(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--LIST--")
	list, err := client.ListNstcpprofile()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateNstcpprofile(*resource)

	// assert.NoError(t, err)

	log.Print("--RENAME--")
	err = client.RenameNstcpprofile(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	log.Print("--DELETE--")
	err = client.DeleteNstcpprofile(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
