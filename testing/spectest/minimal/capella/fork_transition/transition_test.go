package fork_transition

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/capella/fork"
)

func TestMinimal_Capella_Transition(t *testing.T) {
	fork.RunForkTransitionTest(t, "minimal")
}
