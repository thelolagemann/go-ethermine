package ethermine

import "encoding/json"

type baseResponse struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

// MinerHistoryResponse is the json response returned by
// the /miner/:miner/history endpoint.
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

// MinerPayoutsResponse is the json response returned by
// the /miner/:miner/payouts endpoint.
type MinerPayoutsResponse []struct {
	Start  int    `json:"start"`
	End    int    `json:"end"`
	Amount int64  `json:"amount"`
	TxHash string `json:"txHash"`
	PaidOn int    `json:"paidOn"`
}

// MinerRoundsResponse is the json response returned by
// the /miner/:miner/rounds endpoint.
type MinerRoundsResponse []struct {
	Block  int   `json:"block"`
	Amount int64 `json:"amount"`
}

// MinerSettingsResponse is the json response returned by
// the /miner/:miner/settings endpoint.
type MinerSettingsResponse struct {
	Monitor   int    `json:"monitor"`
	MinPayout int64  `json:"minPayout"`
	Email     string `json:"email"`
	IP        string `json:"ip"`
}

// MinerStatisticsResponse is the json response returned by
// the /miner/:miner/currentStats.
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

// PoolMinedBlocksResponse is the json response returned by
// the /blocks/history endpoint.
type PoolMinedBlocksResponse []struct {
	Time       int   `json:"time"`
	NbrBlocks  int   `json:"nbrBlocks"`
	Difficulty int64 `json:"difficulty"`
}

// PoolNetworkStatsResponse is the json response returned by
// the /networkStats endpoint.
type PoolNetworkStatsResponse struct {
	Time       int     `json:"time"`
	BlockTime  float64 `json:"blockTime"`
	Difficulty int64   `json:"difficulty"`
	Hashrate   int64   `json:"hashrate"`
	USD        float64 `json:"usd"`
	BTC        float64 `json:"btc"`
}

// PoolServerHashrateResponse is the json response returned by
// the /servers/history endpoint.
type PoolServerHashrateResponse []struct {
	Time     int     `json:"time"`
	Server   string  `json:"server"`
	Hashrate float64 `json:"hashrate"`
}

// PoolStatsResponse is the json response returned by
// the /poolStats endpoint.
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

// WorkersStatsResponse is the json response returned by
// the various worker endpoints. TODO split?
type WorkerStatsResponse struct {
	Worker           string  `json:"worker"`
	Time             int     `json:"time"`
	LastSeen         int     `json:"lastSeen"`
	ReportedHashrate int     `json:"reportedHashrate"`
	CurrentHashrate  float64 `json:"currentHashrate"`
	ValidShares      int     `json:"validShares"`
	InvalidShares    int     `json:"invalidShares"`
	StaleShares      int     `json:"staleShares"`
	AverageHashrate  float64 `json:"averageHashrate"`
}
