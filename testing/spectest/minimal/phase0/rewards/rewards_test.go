package rewards

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/phase0/rewards"
)

func TestMinimal_Phase0_Rewards(t *testing.T) {
	rewards.RunPrecomputeRewardsAndPenaltiesTests(t, "minimal")
}
