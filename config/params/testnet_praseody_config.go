package params

import "math"

// UsePraseOdyNetworkConfig uses the PraseOdy beacon chain specific network config.
func UsePraseOdyNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{
		// lighthouse
		"enr:-Le4QHeeSdQ-bw7-E9ts1cRkSgmZ5Jn_kZwNz2d6lgvj-dWmN87YR4BAf5uzedTc6MRsDNyCAuOs5oj3GpA-tUGpugkDhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISAjDkLg2lwNpAqAQT4wBJYOgAAAAAAAAABiXNlY3AyNTZrMaECN6t9SCWYDli82RCsM3IGdPOBQ1PXYKwc-0_X1vb_dW6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QOn1Xxt61bjbzQj3_SR4v6XhATbnzfpHzprodQSHKPkpGqtRz6iNEnEfYBpbLCwvf8rcjl8zWOJrPEu6FLZ1-HcBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISn6-WZg2lwNpAqAQT4wBJZpQAAAAAAAAABiXNlY3AyNTZrMaEDeoIwLH0EgoBK-rvWymfZoMfVvuakz_hYKaDS4vrhyJ6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QIniYpAcd4w1M682s3tF7F_zqjEg4nRu8lJXELBCQNQrQOoP1fXRfuCICj51eCQJmSf3EW6d8w0tfoh2OiJwlAIBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQz3gyzg2lwNpAmB1MAAgUCAAAAAAAAADnYiXNlY3AyNTZrMaECCjcCM_ie13QIqkAg12TJlFC6C9_d0todebg-ahRalKqDdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QDo2Sn5g4w_dN1VASUrDfcrUZoz_2PH4R5ITFAZ_3OuMF2GtEtvqoiJgGjbZCYQ6cZBo_917id-1eiaMPD2TX9gBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQFobgAg2lwNpAqAQT_APCN3QAAAAAAAAABiXNlY3AyNTZrMaEDdZHAmo6B5fHY1ndSBT-7cu12vidf1CkElkdybQktNt6DdWRwgiMohHVkcDaCI4I",

		// lodestar small
		"enr:-KG4QJH7wBX6C4U8Ia9s2BUzF_v5bKyTRJkrh9SdMUQ8f1w5RjddgJU_VnIOzgo23CSRdVkEtxfqBAwZUgZUqkfYhz8EgmlkgnY0gmlwhDPeDLODaXA2kCYHUwACBQIAAAAAAAAAOdiJc2VjcDI1NmsxoQLRKq040iFxE6conbY-PrOnZvTZWrNe1fo4YTQYEKBadYN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QJo1hXkNHa5FI5me-8VW2lbnuUi69N_JEa8oqH0Ik1dLFrk2xpgzSGlMuZe2_TKhcBxdGbRnR4TUKfYIvSlZaeEBgmlkgnY0gmlwhAWhuACDaXA2kCoBBP8A8I3dAAAAAAAAAAGJc2VjcDI1NmsxoQO9vb9FwSJh3mxz3wTIdjpBCJr2NMlDnLQ_h3BOOJGQfoN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QNxl-gdB50FYz7MQ9Eyx6kPbH6ssx6vXk8Ppfuv-m_oHE9VrqLVM0vYjyOMAIoAGjZfRTFkbp_CXcAnxAxtnji4BgmlkgnY0gmlwhKfr5ZmDaXA2kCoBBPjAElmlAAAAAAAAAAGJc2VjcDI1NmsxoQLSb0IKr0D_oYSSMbWnRs9uEbQvmgSHq3hkJZ47oeYqPoN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QPsfe0h21rUZa1-Y2LHeK6zHkEW9CgHuvkC-CtXkneFERzeUrKxAcFL0SWg2DJlmjkf4wSkovNwnE3PkzZFiK28BgmlkgnY0gmlwhICMOQuDaXA2kCoBBPjAElg6AAAAAAAAAAGJc2VjcDI1NmsxoQMqVkdXuC1greHkMG4VmTxBHlWeIQ2MLBzPWUH4vZuZEYN1ZHCCIymEdWRwNoIjgw",

		// lodestar big
		"enr:-KG4QMD_1xIgqMBGxuzprUm7SAp5HoFyfjAuc2Y_r6Vmdw6SETb__EtFG09ch-Y9LAXfJxqeLUioYzx8jgZQ2M2Vw2wIgmlkgnY0gmlwhMBjAMGDaXA2kCYHUwAAYDHBAAAAAAAAAAGJc2VjcDI1NmsxoQMXul8Z8ZRhF6KFQmEwB6o0vUjFN1IqUUcQPvR2yQocN4N1ZHCCI4KEdWRwNoIjgw",
		"enr:-KG4QGj23NcTPwhWNkzeTq0PRt6txNsRS0TbAklFDJTkKu59HVkp9vcx2UWv5ANwYDxBxah4NoAMnWk7yGkGTraXnxABgmlkgnY0gmlwhMBjD6eDaXA2kCYHUwAAYECnAAAAAAAAAAGJc2VjcDI1NmsxoQKRzBefFuVAdBTdVhgFk5zUM1szGbG5NsWCv4wQ38iTAYN1ZHCCI4KEdWRwNoIjgw",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// PraseOdyConfig defines the config for the Sepolia beacon chain testnet.
func PraseOdyConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = PraseOdyName

	cfg.MinGenesisTime = 1712319000
	cfg.MinGenesisActiveValidatorCount = 164

	cfg.GenesisDelay = 3600

	cfg.DepositChainID = 30203
	cfg.DepositNetworkID = 30203
	cfg.DepositContractAddress = "0x55155Ca1F57bbDB1e8e10EBB871c00D809E5E84f"

	cfg.GenesisForkVersion = []byte{0x00, 0x32, 0x23, 0x00}
	cfg.AltairForkEpoch = 1
	cfg.AltairForkVersion = []byte{0x00, 0x32, 0x33, 0x00}
	cfg.BellatrixForkEpoch = 10
	cfg.BellatrixForkVersion = []byte{0x00, 0x32, 0x43, 0x00}
	cfg.CapellaForkEpoch = 40
	cfg.CapellaForkVersion = []byte{0x00, 0x32, 0x53, 0x00}
	cfg.DenebForkEpoch = math.MaxUint64
	cfg.DenebForkVersion = []byte{0x00, 0x32, 0x63, 0x00}

	cfg.TerminalTotalDifficulty = "1016000000"

	cfg.SecondsPerSlot = 6
	cfg.SecondsPerETH1Block = 6
	cfg.Eth1FollowDistance = 128

	cfg.BaseRewardFactor = 0
	cfg.BasePenaltyFactor = 0
	cfg.InactivityPenaltyQuotientBellatrix = 1 << 24
	cfg.MaxExcessBalance = 32 * 1e10

	cfg.InitializeForkSchedule()
	return cfg
}
