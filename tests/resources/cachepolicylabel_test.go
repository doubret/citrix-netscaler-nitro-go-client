package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCachepolicylabel(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := Setup_cachepolicylabel(t, client)

	if resource == nil {
		return
	}

	key := resource.ToKey()

	log.Print("--ADD--")
	err := client.AddCachepolicylabel(*resource)

	assert.NoError(t, err)

	log.Print("--LIST--")
	list, err := client.ListCachepolicylabel()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	log.Print("--GET--")
	res, err := client.GetCachepolicylabel(key)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	log.Print("--COUNT--")
	count, err := client.CountCachepolicylabel()

	assert.NoError(t, err)
	assert.NotZero(t, count)

	log.Print("--EXISTS--")
	exists, err := client.ExistsCachepolicylabel(key)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	log.Print("--DELETE--")
	err = client.DeleteCachepolicylabel(key)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
