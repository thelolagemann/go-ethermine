package ethermine

import "encoding/json"

type baseResponse struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type MinerHistoryResponse []struct {
	Time             int     `json:"time"`
	ReportedHashrate int     `json:"reportedHashrate"`
	CurrentHashrate  float64 `json:"currentHashrate"`
	ValidShares      int     `json:"validShares"`
	InvalidShares    int     `json:"invalidShares"`
	StaleShares      int     `json:"staleShares"`
	AverageHashrate  float64 `json:"averageHashrate"`
	ActiveWorkers    int     `json:"activeWorkers"`
}

type MinerPayoutsResponse []struct {
	Start  int    `json:"start"`
	End    int    `json:"end"`
	Amount int64  `json:"amount"`
	TxHash string `json:"txHash"`
	PaidOn int    `json:"paidOn"`
}

type MinerRoundsResponse []struct {
	Block  int   `json:"block"`
	Amount int64 `json:"amount"`
}

type MinerSettingsResponse struct {
	Monitor   int    `json:"monitor"`
	MinPayout int64  `json:"minPayout"`
	Email     string `json:"email"`
	IP        string `json:"ip"`
}

type MinerStatisticsResponse struct {
	Time             int     `json:"time"`
	LastSeen         int     `json:"lastSeen"`
	ReportedHashrate int     `json:"reportedHashrate"`
	CurrentHashrate  float64 `json:"currentHashrate"`
	ValidShares      int     `json:"validShares"`
	InvalidShares    int     `json:"invalidShares"`
	StaleShares      int     `json:"staleShares"`
	AverageHashrate  float64 `json:"averageHashrate"`
	ActiveWorkers    int     `json:"activeWorkers"`
	Unpaid           int64   `json:"unpaid"`
	Unconfirmed      float64 `json:"unconfirmed"`
	CoinsPerMin      float64 `json:"coinsPerMin"`
	UsdPerMin        float64 `json:"usdPerMin"`
	BtcPerMin        float64 `json:"btcPerMin"`
}

type PoolMinedBlocksResponse []struct {
	Time       int   `json:"time"`
	NbrBlocks  int   `json:"nbrBlocks"`
	Difficulty int64 `json:"difficulty"`
}

type PoolNetworkStatsResponse struct {
	Time       int     `json:"time"`
	BlockTime  float64 `json:"blockTime"`
	Difficulty int64   `json:"difficulty"`
	Hashrate   int64   `json:"hashrate"`
	USD        float64 `json:"usd"`
	BTC        float64 `json:"btc"`
}

type PoolServerHashrateResponse []struct {
	Time     int     `json:"time"`
	Server   string  `json:"server"`
	Hashrate float64 `json:"hashrate"`
}

type PoolStatsResponse struct {
	TopMiners   []interface{} `json:"topMiners"`
	MinedBlocks []struct {
		Number int    `json:"number"`
		Miner  string `json:"miner"`
		Time   int    `json:"time"`
	} `json:"minedBlocks"`
	PoolStats struct {
		HashRate      float64 `json:"hashRate"`
		Miners        int     `json:"miners"`
		Workers       int     `json:"workers"`
		BlocksPerHour float64 `json:"blocksPerHour"`
	} `json:"poolStats"`
	Price struct {
		USD float64 `json:"usd"`
		BTC float64 `json:"btc"`
	} `json:"price"`
}
