package ethermine

import (
	"os"
	"testing"
)

// TODO mock

var (
	miner = NewMiner(os.Getenv("ETHERMINE_ADDRESS"), Ethermine)
)

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
