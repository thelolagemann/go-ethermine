package ethermine

import (
	"fmt"
)

// Miner is a convenience struct that embeds a Pool
// alongside a miners address.
type Miner struct {
	address string
	pool    *Pool
}

// Worker is a convenience struct that embeds a Miner
// alongside a workers name, this allows for easy access
// to per-worker endpoints.
type Worker struct {
	*Miner
	name string
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

// Rounds returns a historic record of the miners
// rounds and the base units allocated to them, ordered
// by block number descending.
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

// WorkersStats returns all of the miners workers current statistics,
// ordered by name ascending.
func (m *Miner) WorkersStats() (res []*WorkerStatsResponse, err error) {
	return res, m.get("workers", &res)
}

// WorkersMonitor returns a list of the workers currently active
// under the miner. Note this is for monitoring purposes only and does
// not provide historical data. Dead workers will be displayed for
// up to 7 days.
func (m *Miner) WorkersMonitor() (res []*WorkerStatsResponse, err error) {
	return res, m.get("workers/monitor", res)
}

// Worker is a convenience function for accessing workers. All it does
// is return a Worker embedded with the recieving miner and the workers name.
func (m *Miner) Worker(worker string) *Worker {
	return &Worker{m, worker}
}

// History returns a historic record of the workers statistics,
// ordered by time ascending.
func (w *Worker) History() (res []*WorkerStatsResponse, err error) {
	return res, w.get(fmt.Sprintf("worker/%v/history", w.name), res)
}

// Stats returns the workers current statistics.
func (w *Worker) Stats() (res *WorkerStatsResponse, err error) {
	return res, w.get(fmt.Sprintf("worker/%v/currentStats", w.name), &res)
}
