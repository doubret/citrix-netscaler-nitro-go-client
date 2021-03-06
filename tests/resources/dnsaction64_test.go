package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDnsaction64(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_dnsaction64(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddDnsaction64(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListDnsaction64()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetDnsaction64(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--UPDATE--")
	err = client.UpdateDnsaction64(resource.ToUpdate())

	assert.NoError(t, err)

	log.Print("--COUNT--")
	count, err := client.CountDnsaction64()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsDnsaction64(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteDnsaction64(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
