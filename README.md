# Vero

[![Go Reference](https://pkg.go.dev/badge/github.com/pastc/vero.svg)](https://pkg.go.dev/github.com/pastc/vero)

The vero package generates random numbers for online casinos games such as dice, roll, crash, etc.

All the algorithms are provably fair.

## Install

```shell
go get github.com/pastc/vero
```

## Guide

### Crash

###### Controllable variables

```go
// HouseEdge i.e, percentage that the house gets
HouseEdge = 6.66
```

###### Explanation

The server should first generate a chain of 10 million SHA256 hashes, starting with a server secret that has been
repeatedly fed the output of SHA256 back into itself 10 million times. Then, the crash game is played through this chain
of hashes in reverse order, using the values as source data for generating each game's outcome.

Anyone can easily verify the integrity of the whole chain as the server should publish the hash used to calculate the
outcome after each game ends.

If you apply the SHA256 function to a revealed seed, you'll get the hash for the previously played game, and so on until
you get the hash for the first ever played game round on the chain.

Though, for security reasons, it would not be safe to keep using the same hash chain for a lot of games (> 1 million) in
a row, as these games would last several years. The longer the chain is in use, the higher the risk is that the
originating seed is found by a malicious third-party (though it is unlikely). Thus, the server should periodically
update the base seed to ensure that the lifetime of a chain is not too long and the risk is minimised. When the server
updates the chain, they should publish the first hash and the corresponding matchId. By publishing the seeds, the server
is preventing itself from switching or modifing the chain.

### Dice

###### Controllable variables

```go
None
```

###### Explanation

### Roll

###### Controllable variables

```go
// Maximum is the maximum value that can be rolled
var Maximum = 15

// ColorMap is colors mapped to values
var ColorMap = map[int]string{
0:  "Green",
1:  "Red",
2:  "Red",
3:  "Red",
4:  "Red",
5:  "Red",
6:  "Red",
7:  "Red",
8:  "Black",
9:  "Black",
10: "Black",
11: "Black",
12: "Black",
13: "Black",
14: "Black",
}

// BaitMap is baits mapped to values
var BaitMap = map[int]string{
4:  "Bait",
11: "Bait",
}
```

## Documentation

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here:
https://pkg.go.dev/github.com/pastc/vero
