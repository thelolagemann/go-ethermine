# go-ethermine

> An unoffical golang client for ethermine and their generic pool API.

## Getting started

### Install

```
go get -u github.com/thelolagemann/go-ethermine
```

### Example usage

```golang
...
stats, err := Ethermine.Stats()
if err != nil {
    // handle err...
}
fmt.Printf("Ethermine is mining $%.2f an hour!", (stats.PoolStats.BlocksPerHour*5)*stats.Price.USD)
// Output: Ethermine is mining $430049.51 an hour!
...
```
