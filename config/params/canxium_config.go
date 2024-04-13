package params

import "math"

// UseCanxiumNetworkConfig uses the Canxium beacon chain specific network config.
func UseCanxiumNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 3459834
	cfg.BootstrapNodes = []string{
		// Lighhouse bootnodes
		"enr:-Le4QHeeSdQ-bw7-E9ts1cRkSgmZ5Jn_kZwNz2d6lgvj-dWmN87YR4BAf5uzedTc6MRsDNyCAuOs5oj3GpA-tUGpugkDhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISAjDkLg2lwNpAqAQT4wBJYOgAAAAAAAAABiXNlY3AyNTZrMaECN6t9SCWYDli82RCsM3IGdPOBQ1PXYKwc-0_X1vb_dW6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QOn1Xxt61bjbzQj3_SR4v6XhATbnzfpHzprodQSHKPkpGqtRz6iNEnEfYBpbLCwvf8rcjl8zWOJrPEu6FLZ1-HcBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISn6-WZg2lwNpAqAQT4wBJZpQAAAAAAAAABiXNlY3AyNTZrMaEDeoIwLH0EgoBK-rvWymfZoMfVvuakz_hYKaDS4vrhyJ6DdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QIniYpAcd4w1M682s3tF7F_zqjEg4nRu8lJXELBCQNQrQOoP1fXRfuCICj51eCQJmSf3EW6d8w0tfoh2OiJwlAIBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQz3gyzg2lwNpAmB1MAAgUCAAAAAAAAADnYiXNlY3AyNTZrMaECCjcCM_ie13QIqkAg12TJlFC6C9_d0todebg-ahRalKqDdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QDo2Sn5g4w_dN1VASUrDfcrUZoz_2PH4R5ITFAZ_3OuMF2GtEtvqoiJgGjbZCYQ6cZBo_917id-1eiaMPD2TX9gBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcIQFobgAg2lwNpAqAQT_APCN3QAAAAAAAAABiXNlY3AyNTZrMaEDdZHAmo6B5fHY1ndSBT-7cu12vidf1CkElkdybQktNt6DdWRwgiMohHVkcDaCI4I",

		// lodestar bootnodes
		"enr:-KG4QJH7wBX6C4U8Ia9s2BUzF_v5bKyTRJkrh9SdMUQ8f1w5RjddgJU_VnIOzgo23CSRdVkEtxfqBAwZUgZUqkfYhz8EgmlkgnY0gmlwhDPeDLODaXA2kCYHUwACBQIAAAAAAAAAOdiJc2VjcDI1NmsxoQLRKq040iFxE6conbY-PrOnZvTZWrNe1fo4YTQYEKBadYN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QJo1hXkNHa5FI5me-8VW2lbnuUi69N_JEa8oqH0Ik1dLFrk2xpgzSGlMuZe2_TKhcBxdGbRnR4TUKfYIvSlZaeEBgmlkgnY0gmlwhAWhuACDaXA2kCoBBP8A8I3dAAAAAAAAAAGJc2VjcDI1NmsxoQO9vb9FwSJh3mxz3wTIdjpBCJr2NMlDnLQ_h3BOOJGQfoN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QNxl-gdB50FYz7MQ9Eyx6kPbH6ssx6vXk8Ppfuv-m_oHE9VrqLVM0vYjyOMAIoAGjZfRTFkbp_CXcAnxAxtnji4BgmlkgnY0gmlwhKfr5ZmDaXA2kCoBBPjAElmlAAAAAAAAAAGJc2VjcDI1NmsxoQLSb0IKr0D_oYSSMbWnRs9uEbQvmgSHq3hkJZ47oeYqPoN1ZHCCIymEdWRwNoIjgw",
		"enr:-KG4QPsfe0h21rUZa1-Y2LHeK6zHkEW9CgHuvkC-CtXkneFERzeUrKxAcFL0SWg2DJlmjkf4wSkovNwnE3PkzZFiK28BgmlkgnY0gmlwhICMOQuDaXA2kCoBBPjAElg6AAAAAAAAAAGJc2VjcDI1NmsxoQMqVkdXuC1greHkMG4VmTxBHlWeIQ2MLBzPWUH4vZuZEYN1ZHCCIymEdWRwNoIjgw",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// CanxiumConfig defines the config for the Sepolia beacon chain testnet.
func CanxiumConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = CanxiumName

	cfg.MinGenesisTime = 1713571200 // 20/4/2024 00:00:00
	cfg.MinGenesisActiveValidatorCount = 164

	cfg.GenesisDelay = 864000 // 10 days

	cfg.DepositChainID = 3003
	cfg.DepositNetworkID = 3003
	cfg.DepositContractAddress = "0xE616698D7c13E46C538f380da7b67E8FA7929a24"

	cfg.GenesisForkVersion = []byte{0x00, 0x30, 0x03, 0x00}
	cfg.AltairForkVersion = []byte{0x00, 0x30, 0x13, 0x00}
	cfg.AltairForkEpoch = 16
	cfg.BellatrixForkVersion = []byte{0x00, 0x30, 0x23, 0x00}
	cfg.BellatrixForkEpoch = 450 // Estimate 1/5/2024
	cfg.CapellaForkVersion = []byte{0x00, 0x30, 0x33, 0x00}
	cfg.CapellaForkEpoch = 27450 // Estimate 30/6/2024
	cfg.DenebForkVersion = []byte{0x00, 0x30, 0x43, 0x00}
	cfg.DenebForkEpoch = math.MaxUint64

	cfg.TerminalTotalDifficulty = "79000000000000000000"

	cfg.SecondsPerSlot = 6
	cfg.SecondsPerETH1Block = 6
	cfg.Eth1FollowDistance = 2048

	cfg.BaseRewardFactor = 0
	cfg.BasePenaltyFactor = 128
	cfg.InactivityPenaltyQuotientBellatrix = 1 << 22
	cfg.MaxExcessBalance = 32 * 1e10

	cfg.InitializeForkSchedule()
	return cfg
}
