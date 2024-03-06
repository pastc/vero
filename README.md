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

###### Example

```go
seed := "2826d440b0fcad643e3008693c3a93ef81b31675ca00d686e44c40d5e83d7bb6"

crashPoint, err := vero.Crash(seed)
if err != nil {
  log.Fatal(err)
}

// crashPoint should be 126
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

###### Example

```go
serverSeed := "1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d"
clientSeed := "5b60f37f764fdb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71"
nonce := 1

value, err := vero.Dice(serverSeed, clientSeed, nonce, 0)
if err != nil {
  log.Fatal(err)
}

// value should be 7473
```

###### Explanation

The server seed is generated first before you specify your client seed. Both seeds together prevent manipulation from
the server and verifies the roll integrity after the result calculation. Every roll has a unique server seed randomly
generated in advance, this server seed will only be updated when you choose to update your client seed. The server
hashes the server seed with the SHA256 cryptographic function and then publishes the hashed server seeds for the player
to see.

Due to this applied hashing function, the player can verify the integrity of the roll by updating your client seed to
get the server seed of the player, which they then can apply the function to get the roll number. The client seed can be
edited freely by users before each roll.

As the client seed affects every roll result, changing it to any seed of your choice at any time means you can ensure
that it's impossible for the server to manipulate the result. However, the SHA512 function we use to generate the roll
is deterministic, if the client seed is combined with the same server seed, it will generate exactly the same roll
result every time. This could be used to abuse the system, so we use something called a 'nonce' which prevents this from
being abusable. Each roll done using the same server seed & client seed pair will also be paired with a different nonce,
which is simply a number starting at 0 and incremented by 1 for each roll done.

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

###### Example

```go
serverSeed := "1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d"
clientSeed := "1c064b20e2ed52a5c4db0361a2523e8901db2342f95bd0dd1d9a68a46b8cc483"
nonce := 5345510

color, value, err := vero.Dice(serverSeed, clientSeed, nonce, 0)
if err != nil {
  log.Fatal(err)
}

// color, value should be Red, 1
```

###### Explanation

### Plinko

###### Controllable variables

```go
None
```

###### Example

```go
serverSeed := "1c5cff3922c8dc1fc9188b3cc2805acdafb6b3a51f51860b59f98eb1753c170d"
clientSeed := "5b60f37f764fdb9700d202d6caf3a0cf1d5e67020b0ce1f6570d16f34150cc71"
nonce := 493587

value, err := vero.Plinko(serverSeed, clientSeed, nonce, 0)
if err != nil {
  log.Fatal(err)
}

// value should be -1 (left)
```

###### Explanation

## Documentation

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here:
https://pkg.go.dev/github.com/pastc/vero
