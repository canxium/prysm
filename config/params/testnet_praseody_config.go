package params

import (
	"math"
)

// UsePraseOdyNetworkConfig uses the PraseOdy beacon chain specific network config.
func UsePraseOdyNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{
		"enr:-Le4QHeeSdQ-bw7-E9ts1cRkSgmZ5Jn_kZwNz2d6lgvj-dWmN87YR4BAf5uzedTc6MRsDNyCAuOs5oj3GpA-tUGpugkDhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISAjDkLg2lwNpAqAQT4wBJYOgAAAAAAAAABiXNlY3AyNTZrMaECN6t9SCWYDli82RCsM3IGdPOBQ1PXYKwc-0_X1vb_dW6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QOn1Xxt61bjbzQj3_SR4v6XhATbnzfpHzprodQSHKPkpGqtRz6iNEnEfYBpbLCwvf8rcjl8zWOJrPEu6FLZ1-HcBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISn6-WZg2lwNpAqAQT4wBJZpQAAAAAAAAABiXNlY3AyNTZrMaEDeoIwLH0EgoBK-rvWymfZoMfVvuakz_hYKaDS4vrhyJ6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QCRBiT1G2typzKaccAOndvpCIPaPxkGj4uoNkjjTKWR0Lnd7ryfaqoT8Q1kVupt-GgjUuIy9EQ-Pmb_7noI7yR4ChGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQz3gyzg2lwNpAmB1MAAgUCAAAAAAAAADnYiXNlY3AyNTZrMaED7RB_J71JjsurCW-DvKGPoPQmm92bXlogSKeUlHey1o-DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QDo2Sn5g4w_dN1VASUrDfcrUZoz_2PH4R5ITFAZ_3OuMF2GtEtvqoiJgGjbZCYQ6cZBo_917id-1eiaMPD2TX9gBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQFobgAg2lwNpAqAQT_APCN3QAAAAAAAAABiXNlY3AyNTZrMaEDdZHAmo6B5fHY1ndSBT-7cu12vidf1CkElkdybQktNt6DdWRwgiMohHVkcDaCI4I",
		"enr:-KG4QJH7wBX6C4U8Ia9s2BUzF_v5bKyTRJkrh9SdMUQ8f1w5RjddgJU_VnIOzgo23CSRdVkEtxfqBAwZUgZUqkfYhz8EgmlkgnY0gmlwhDPeDLODaXA2kCYHUwACBQIAAAAAAAAAOdiJc2VjcDI1NmsxoQLRKq040iFxE6conbY-PrOnZvTZWrNe1fo4YTQYEKBadYN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QJo1hXkNHa5FI5me-8VW2lbnuUi69N_JEa8oqH0Ik1dLFrk2xpgzSGlMuZe2_TKhcBxdGbRnR4TUKfYIvSlZaeEBgmlkgnY0gmlwhAWhuACDaXA2kCoBBP8A8I3dAAAAAAAAAAGJc2VjcDI1NmsxoQO9vb9FwSJh3mxz3wTIdjpBCJr2NMlDnLQ_h3BOOJGQfoN1ZHCCIymEdWRwNoIjgw",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PraseOdyConfig defines the config for the Sepolia beacon chain testnet.
func PraseOdyConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.MinGenesisTime = 1711616400
	cfg.GenesisDelay = 7200
	cfg.MinGenesisActiveValidatorCount = 164
	cfg.ConfigName = PraseOdyName
	cfg.GenesisForkVersion = []byte{0x00, 0x32, 0x23, 0x00}
	cfg.SecondsPerETH1Block = 8
	cfg.DepositChainID = 30203
	cfg.DepositNetworkID = 30203
	cfg.AltairForkEpoch = 4
	cfg.AltairForkVersion = []byte{0x00, 0x32, 0x33, 0x00}
	cfg.BellatrixForkEpoch = 10
	cfg.BellatrixForkVersion = []byte{0x00, 0x32, 0x43, 0x00}
	cfg.CapellaForkEpoch = 225
	cfg.CapellaForkVersion = []byte{0x00, 0x32, 0x53, 0x00}
	cfg.DenebForkEpoch = math.MaxUint64
	cfg.DenebForkVersion = []byte{0x00, 0x32, 0x63, 0x00}
	cfg.TerminalTotalDifficulty = "4065600000"
	cfg.DepositContractAddress = "0x55155Ca1F57bbDB1e8e10EBB871c00D809E5E84f"
	cfg.SecondsPerSlot = 6
	cfg.EjectionBalance = 160 * 1e9
	cfg.MaxEffectiveBalance = 320 * 1e9
	cfg.BaseRewardFactor = 0
	cfg.InitializeForkSchedule()
	return cfg
}
