package ethermine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

// TODO mock

var (
	miner = NewMiner(os.Getenv("ETHERMINE_ADDRESS"), Ethermine)

	mockResults map[string]json.RawMessage
)

type mockClient struct{}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	if endpoint := req.URL.Path; endpoint != "" {
		fmt.Println(endpoint)
		if val, ok := mockResults[endpoint]; ok {
			r := ioutil.NopCloser(bytes.NewReader(val))
			return func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			}(req)
		}
	}
	return nil, errors.New("couldn't mock")
}

func init() {
	if os.Getenv("ETHERMINE_ADDRESS") == "" {
		client = &mockClient{}
	}

	b, err := os.ReadFile("responses.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&mockResults); err != nil {
		panic(err)
	}
}

func TestMiner_History(t *testing.T) {
	if _, err := miner.History(); err != nil {
		t.Error(err)
	}
}

func TestMiner_Payouts(t *testing.T) {
	if _, err := miner.Payouts(); err != nil {
		t.Error(err)
	}
}

func TestMiner_Rounds(t *testing.T) {
	if _, err := miner.Rounds(); err != nil {
		t.Error(err)
	}
}

func TestMiner_Settings(t *testing.T) {
	if _, err := miner.Settings(); err != nil {
		t.Error(err)
	}
}

func TestMiner_Stats(t *testing.T) {
	if _, err := miner.Stats(); err != nil {
		t.Error(err)
	}
}

func TestMiner_WorkerHistory(t *testing.T) {
	if _, err := miner.Worker("miner").History(); err != nil {
		t.Error(err)
	}
}

func TestMiner_WorkerStats(t *testing.T) {
	if _, err := miner.Worker("miner").Stats(); err != nil {
		t.Error(err)
	}
}

func TestMiner_WorkersMonitor(t *testing.T) {
	if _, err := miner.WorkersMonitor(); err != nil {
		t.Error(err)
	}
}

func TestMiner_WorkersStats(t *testing.T) {
	if _, err := miner.WorkersStats(); err != nil {
		t.Error(err)
	}
}

func TestPool_MinedBlocks(t *testing.T) {
	if _, err := Ethermine.MinedBlocks(); err != nil {
		t.Error(err)
	}
}

func TestPool_NetworkStats(t *testing.T) {
	if _, err := Ethermine.NetworkStats(); err != nil {
		t.Error(err)
	}
}

func TestPool_ServerHashrate(t *testing.T) {
	if _, err := Ethermine.ServerHashrate(); err != nil {
		t.Error(err)
	}
}

func TestPool_Stats(t *testing.T) {
	if _, err := Ethermine.Stats(); err != nil {
		t.Error(err)
	}
}
