package epoch_processing

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/deneb/epoch_processing"
)

func TestMinimal_Deneb_EpochProcessing_JustificationAndFinalization(t *testing.T) {
	epoch_processing.RunJustificationAndFinalizationTests(t, "minimal")
}
