package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestServicegroupServicegroupmemberBinding(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_servicegroup_servicegroupmember_binding(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--COUNT--")
	count0, err := client.CountServicegroupServicegroupmemberBinding()

	assert.NoError(t, err)

	log.Print("--ADD--")
	err = client.AddServicegroupServicegroupmemberBinding(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListServicegroupServicegroupmemberBinding()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--COUNT--")
	count, err := client.CountServicegroupServicegroupmemberBinding()

	assert.NoError(t, err)
	assert.NotZero(t, count)
	assert.Equal(t, count, count0+1)

	log.Print("--GET--")
	res, err := client.GetServicegroupServicegroupmemberBinding(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--EXISTS--")
	exists, err := client.ExistsServicegroupServicegroupmemberBinding(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteServicegroupServicegroupmemberBinding(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
