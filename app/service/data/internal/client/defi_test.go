package client

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_DefiGetTokenPrice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		resp, err := testClient.DefiGetTokenPrice(testCtx, "ETH", "USD")
		assert.Nil(t, err)
		spew.Dump(resp)
	})
}
