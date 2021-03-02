package ethermine

import (
	"fmt"
)

// Miner
type Miner struct {
	address string
	pool    *Pool
}

// NewMiner create a new miner with the provided address
// and pool.
func NewMiner(address string, pool *Pool) *Miner {
	return &Miner{address, pool}
}

func (m *Miner) get(endpoint string, bind interface{}) error {
	return m.pool.get(fmt.Sprintf("miner/%v/%v", m.address, endpoint), bind)
}

// History returns a historic record of the miners
// performance, ordered by time ascending.
func (m *Miner) History() (res *MinerHistoryResponse, err error) {
	return res, m.get("history", &res)
}

// Payouts returns a historic record of the miners
// payouts, ordered by time ascending.
func (m *Miner) Payouts() (res *MinerPayoutsResponse, err error) {
	return res, m.get("payouts", &res)
}

// Rounds... TODO
func (m *Miner) Rounds() (res *MinerRoundsResponse, err error) {
	return res, m.get("rounds", &res)
}

// Settings returns the miners current settings. Both the
// email and ip fields will be masked.
func (m *Miner) Settings() (res *MinerSettingsResponse, err error) {
	return res, m.get("settings", &res)
}

// Stats returns a historic record of the miners stats,
// ordered by time ascending.
func (m *Miner) Stats() (res *MinerStatisticsResponse, err error) {
	return res, m.get("currentStats", &res)
}
