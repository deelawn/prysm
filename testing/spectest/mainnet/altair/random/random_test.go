package random

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v2/config/features"
	"github.com/prysmaticlabs/prysm/v2/testing/spectest/shared/altair/sanity"
)

func TestMain(m *testing.M) {
	resetCfg := features.InitWithReset(&features.Flags{EnableBalanceTrieComputation: true})
	defer resetCfg()
	m.Run()
}

func TestMainnet_Altair_Random(t *testing.T) {
	sanity.RunBlockProcessingTest(t, "mainnet", "random/random/pyspec_tests")
}
