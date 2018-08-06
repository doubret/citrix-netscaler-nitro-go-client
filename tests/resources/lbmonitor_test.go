package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLbmonitor(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_lbmonitor(t, client)

	if resource == nil {
		return
	}

	log.Print("--ADD--")
	err := client.AddLbmonitor(*resource)

	assert.NoError(t, err)

	if tearDown != nil {
		tearDown()
	}
}
