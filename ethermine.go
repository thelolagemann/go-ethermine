// Package ethermine is an unofficial golang client for ethpool,
// ethermine and flypool, using their generic pool API. View
// https://api.ethermine.org/docs for detailed documentation
// of their API.
package ethermine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	client httpClient = &http.Client{}
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Pool is the base client used for sending requests to
// the generic API.
type Pool struct {
	url string
}

var (
	// Ethermine default client for Ethermine
	Ethermine = &Pool{"https://api.ethermine.org"}

	// EthermineETC default client for ETC ethermine
	EthermineETC = &Pool{"https://api-etc.ethermine.org"}

	// Flypool default client for flypool
	Flypool = &Pool{"https://api.flypool.org"}

	// Ethpool default client for Ethpool
	Ethpool = &Pool{"https://api.ethpool.org"}
)

func (p *Pool) get(endpoint string, bind interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", p.url, endpoint), nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%v returned non-ok http code, %v", endpoint, res.StatusCode)
	}

	if bind != nil {
		var decode baseResponse

		if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {
			return fmt.Errorf("error decoding json: %v", err)
		}

		if err := json.Unmarshal(decode.Data, &bind); err != nil {
			return fmt.Errorf("couldn't unmarshal json: %v", err)
		}

	}

	return nil
}

// MinedBlocks returns a list of blocks mined by the pool,
// ordered by time ascending.
func (p *Pool) MinedBlocks() (pRes *PoolMinedBlocksResponse, err error) {
	return pRes, p.get("blocks/history", &pRes)
}

// Miner returns a Miner with the recieving pool embedded.
func (p *Pool) Miner(address string) *Miner {
	return &Miner{address, p}
}

// NetworkStats returns current network stats, including
// hashrate and difficulty.
func (p *Pool) NetworkStats() (pRes *PoolNetworkStatsResponse, err error) {
	return pRes, p.get("networkStats", &pRes)
}

// ServerHashrate returns a historic record of the pools
// hashrate, ordered by time ascending.
func (p *Pool) ServerHashrate() (pRes *PoolServerHashrateResponse, err error) {
	return pRes, p.get("servers/history", &pRes)
}

// Stats returns basic pool stats such as recently mined
// blocks, current pricing information, hashrate and more.
func (p *Pool) Stats() (pRes *PoolStatsResponse, err error) {
	return pRes, p.get("poolStats", &pRes)
}
