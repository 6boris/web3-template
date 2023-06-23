package client

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_DefiGetPrice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data, err := testClient.DefiGetPrice(testCtx, "ETH", "USD")
		assert.Nil(t, err)
		spew.Dump(data)
	})
}
