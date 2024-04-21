//go:build !noMainnetGenesis
// +build !noMainnetGenesis

package genesis

import (
	_ "embed"

	"github.com/prysmaticlabs/prysm/v5/config/params"
)

var (
	//go:embed mainnet.ssz.snappy
	mainnetRawSSZCompressed []byte // 1.8Mb
	//go:embed praseody.ssz.snappy
	praseodyRawSSZCompressed []byte // 1.8Mb
	//go:embed canxium.ssz.snappy
	canxiumRawSSZCompressed []byte // 1.8Mb
)

func init() {
	embeddedStates[params.MainnetName] = &mainnetRawSSZCompressed
	embeddedStates[params.PraseOdyName] = &praseodyRawSSZCompressed
	embeddedStates[params.CanxiumName] = &canxiumRawSSZCompressed
}
