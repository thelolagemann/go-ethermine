# go-ethermine

[![GoDoc reference](https://img.shields.io/badge/docs-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/thelolagemann/go-ethermine) [![GitHub](https://img.shields.io/github/license/thelolagemann/go-ethermine?style=flat-square)](https://github.com/thelolagemann/go-ethermine/blob/main/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thelolagemann/go-ethermine?style=flat-square&l)

> An unoffical golang client for ethermine and their generic pool API.

## Getting started

### Install

```
go get -u github.com/thelolagemann/go-ethermine
```

### Example usage

```golang
...
// pool interaction
stats, err := Ethermine.Stats()
if err != nil {
    // handle err...
}
fmt.Printf("Ethermine is mining $%.2f an hour!", (stats.PoolStats.BlocksPerHour*5)*stats.Price.USD)
// Output: Ethermine is mining $430049.51 an hour!
...
// miner interaction
miner := Ethermine.Miner("ADDRESS")
stats, err := miner.Stats()
if err != nil {
    // handle err
}
fmt.Printf("Unpaid balance: %v", stats.Unpaid)
...
// worker interaction
worker := miner.Worker("MINER_NAME")
stats, err := worker.Stats()
if err !=- nil {
    // handle err`
}
fmt.Printf("Worker hashrate: %v", stats.CurrentHashrate)
```

View the godocs for more detailed documentation.
