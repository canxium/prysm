package operations

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/capella/operations"
)

func TestMainnet_Capella_Operations_VoluntaryExit(t *testing.T) {
	operations.RunVoluntaryExitTest(t, "mainnet")
}
