package params

import "math"

// UseCanxiumNetworkConfig uses the Canxium beacon chain specific network config.
func UseCanxiumNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 3459834
	cfg.BootstrapNodes = []string{
		// lighthouse
		"enr:-KG4QOOxrHySRX2492UCkBI0SaaXKagaiKT5EA_eWqRCRqNBOCG-uUL3L4f23hld56TGJxLKKvje4wVMmjMWldT_HkkBgmlkgnY0gmlwhAWhuACDaXA2kCoBBP8A8I3dAAAAAAAAAAGJc2VjcDI1NmsxoQMEhQ0b95Na4X8HKPPiXkOfBFjtrwbFfxv2zQd9ftEjDoN1ZHCCIyiEdWRwNoIjgg", // 40
		"enr:-KG4QP-8mQn4mF7w0sEgsppcqQKC8OajMziibw5Nwdx6aS7XJKU37AaLRLIK2_SZKw-YN-RxCTPFrB_NV6kW4qOqdwQBgmlkgnY0gmlwhDENjdaDaXA2kCoBBPjAEjvJAAAAAAAAAAGJc2VjcDI1NmsxoQM3hKBApmcdjSYlc0v-dfz5Ayxo8iUQhkIO-WlPTJ6OIIN1ZHCCIyiEdWRwNoIjgg", // 14
		"enr:-KG4QDdnAQ0BXczORLH-NMhOmgJOwByQMlaFe89ZE6bKBxDCMKzYt4W8oh3jUTSUu4srZx_IDBja85jrM1Uc-j4HN1IBgmlkgnY0gmlwhLIQidGDaXA2kCoCR4AAEiOnAAAAAAAAAAGJc2VjcDI1NmsxoQPkMcI-M6EAYP6ZhbZJ3qndpqRpFMQXHZk6JKdVp8L3coN1ZHCCIyiEdWRwNoIjgg", // 09

		// lodestar bootnodes
		"enr:-KG4QMC-zvtTtzzKxeIh79dU43e8lB0mmyAi5LGGx55S3fulSY5NLtuRk-5XREhxWXCIgui659j6ybm6RoLYkUxAam8DgmlkgnY0gmlwhAWhuACDaXA2kCoBBP8A8I3dAAAAAAAAAAGJc2VjcDI1NmsxoQMnWOr_eBqyHfkFl38LbEcHbJUklOIvi-e0g6Lc2voJ9IN1ZHCCIymEdWRwNoIjgw", // 51
		"enr:-KG4QCN5VGM8m82K5eGFrIcfxRFT-fZ07xuNsNQhSNnglZLWeD62wo3yHX4a_3s_NRYZGsoPL5MTNHI4uEgsfRCdYt4BgmlkgnY0gmlwhDENjdaDaXA2kCoBBPjAEjvJAAAAAAAAAAGJc2VjcDI1NmsxoQMKzy96zC5aeVJlTH2iZEhBXC7OsFePTavXie0tPZP1hIN1ZHCCIymEdWRwNoIjgw", // 49
		"enr:-KG4QCqcddZnwxPt9V_Wya3i9pfcdxfBkHoPPmgAZG8mkx-3DLoMwxC-BqaqoICiCXzDotRHw50Abi11_C1fODZs5kgBgmlkgnY0gmlwhLIQidGDaXA2kCoCR4AAEiOnAAAAAAAAAAGJc2VjcDI1NmsxoQJmXcAl4vGAoCB9-k51rvJz8Lmo21HvWe8CEIFa0ezUKoN1ZHCCIymEdWRwNoIjgw", // 17

	}
	OverrideBeaconNetworkConfig(cfg)
}

// CanxiumConfig defines the config for the Sepolia beacon chain testnet.
func CanxiumConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.ConfigName = CanxiumName

	cfg.MinGenesisTime = 1713657600 // Apr 21, 2024, 00:00 AM UTC
	cfg.MinGenesisActiveValidatorCount = 166

	cfg.GenesisDelay = 779671 // 9 days

	cfg.DepositChainID = 3003
	cfg.DepositNetworkID = 3003
	cfg.DepositContractAddress = "0xE616698D7c13E46C538f380da7b67E8FA7929a24"

	cfg.GenesisForkVersion = []byte{0x00, 0x30, 0x03, 0x00}
	cfg.AltairForkVersion = []byte{0x00, 0x30, 0x13, 0x00}
	cfg.AltairForkEpoch = 5
	cfg.BellatrixForkVersion = []byte{0x00, 0x30, 0x23, 0x00}
	cfg.BellatrixForkEpoch = 10 // Estimate Apr 30, 2024
	cfg.CapellaForkVersion = []byte{0x00, 0x30, 0x33, 0x00}
	cfg.CapellaForkEpoch = 45000 // Estimate 8 August 2024
	cfg.DenebForkVersion = []byte{0x00, 0x30, 0x43, 0x00}
	cfg.DenebForkEpoch = math.MaxUint64

	cfg.TerminalTotalDifficulty = "86680000000000000000" // Estimate 20 June 2024 if hashrate = 45TH

	cfg.SecondsPerSlot = 6
	cfg.SecondsPerETH1Block = 6
	cfg.Eth1FollowDistance = 2048

	cfg.BaseRewardFactor = 512
	cfg.InactivityPenaltyQuotientBellatrix = 1 << 22
	cfg.MaxExcessBalance = 32 * 1e10

	cfg.InitializeForkSchedule()
	return cfg
}
