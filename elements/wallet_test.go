package elements_test

import (
	"testing"

	"github.com/rddl-network/elements-rpc/types"
	"github.com/rddl-network/rddl-prometheus-exporter/elements"
	"github.com/stretchr/testify/assert"
)

func TestGetAsset(t *testing.T) {
	casino := types.GetBalanceResult{"bitcoin": 0.1}
	player := types.GetBalanceResult{}

	btcBalance, err := elements.GetAsset(casino)
	assert.Equal(t, btcBalance, 0.1)
	assert.NoError(t, err)

	btcBalance, err = elements.GetAsset(player)
	assert.Equal(t, btcBalance, 0.0)
	assert.Error(t, err)
}
