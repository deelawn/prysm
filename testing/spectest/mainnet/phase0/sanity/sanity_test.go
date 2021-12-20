package sanity

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v2/config/features"
)

func TestMain(m *testing.M) {
	resetCfg := features.InitWithReset(&features.Flags{EnableBalanceTrieComputation: true})
	defer resetCfg()
	m.Run()
}
