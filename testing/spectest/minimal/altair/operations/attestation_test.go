package operations

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/altair/operations"
)

func TestMinimal_Altair_Operations_Attestation(t *testing.T) {
	operations.RunAttestationTest(t, "minimal")
}
