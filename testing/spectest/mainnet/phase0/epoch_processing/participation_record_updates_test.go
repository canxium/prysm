package epoch_processing

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/phase0/epoch_processing"
)

func TestMainnet_Phase0_EpochProcessing_ParticipationRecordUpdates(t *testing.T) {
	epoch_processing.RunParticipationRecordUpdatesTests(t, "mainnet")
}
