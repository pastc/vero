# Vero

[![Go Reference](https://pkg.go.dev/badge/github.com/pastc/vero.svg)](https://pkg.go.dev/github.com/pastc/vero)

The vero package generates random numbers for online casinos games such as dice, roll, crash, etc.

All the algorithms are provably fair.

## Install

```shell
go get github.com/pastc/vero/v2
```

## Features

- No third-party dependencies
- As minimal as possible
- Provably fair
- Many games:
  - Crash
  - Dice
  - Roll
  - Plinko

## Guide

### Common arguments

```go
// Generated automatically. Is private and is changed periodically. It should become public after being decommissioned.
var serverSeed string

// Generated uniquely to each player. Can be changed anytime by the player themselves.
var clientSeed string

// Generated automatically. Is public and is changed periodically.
var publicSeed string

// A number that is incremented by 1 for each game played.
var nonce int

// A number that is incremented by 1 if the randomly generated value goes out of bounds and eventually exhausts all available random numbers within the available hash.
// Should be 0 when calling the function from outside.
var iteration int
```

### Crash

#### Arguments

```go
var serverSeed string
// houseEdge i.e, percentage that the house gets.
var houseEdge float64
```

#### Example

```go
// Remember to divide the crashPoint by 100 to get the percentage
crashPoint, err := vero.Crash(serverSeed, houseEdge)
if err != nil {
  log.Fatal(err)
}
```

#### Explanation

1. The function calculates an HMAC-SHA256 hash using the `serverSeed` and a combined `seed`. This hash is used as a
   source of randomness.

2. The most significant 52 bits of the hash are extracted and interpreted as a hexadecimal number, which is then
   converted to an integer (`h`).

3. The value `e` (2^52) is calculated, which is approximately 4.5035e+15. This value represents the maximum value that
   can be represented precisely in the mantissa of a 64-bit floating-point number.

4. The function then calculates the crash point multiplier (`result`) using the values of `h` and `e`. The
   formula `(100*e - float64(h)) / (e - float64(h))` maps the value of `h` (which is in the range `[0, e]`) to a value
   in the range `[100, infinity]`.

5. The `houseEdgeModifier` is calculated based on the specified house edge percentage. For example, if the house edge is
   5%, the `houseEdgeModifier` will be 0.95 with the lowest crashing point of 100.

6. The final crash point multiplier (`endResult`) is calculated by multiplying result by `houseEdgeModifier` and
   ensuring that it is at least 100 (the minimum crash point).

7. The function returns the crash point multiplier `endResult` as an integer, which represents the crash point
   multiplier.

### Dice

#### Arguments

```go
var serverSeed string
var clientSeed string
var nonce int
var iteration int
```

#### Example

```go
// Remember to divide the value by 100 to get a number from 0 to 99.99
value, err := vero.Dice(serverSeed, clientSeed, nonce, 0)
if err != nil {
  log.Fatal(err)
}
```

#### Explanation

1. The function calculates an HMAC-SHA512 hash using the `serverSeed` and a combined `seed`. This hash is used as a
   source of randomness.

2. The `GetLucky` function extracts a substring of length 5 from the `hash` string starting at the position `index*5`.
   This substring is then converted from a hexadecimal string to an integer. The function returns the random integer and
   any error that may have occurred during the conversion.

3. The for loop ensures that the `lucky` integer is above 10^6 (1000000) since it will be divided by 10^4 (10000) later.

   1. If it is under 10^6, then the `index` is incremented by 1 and the `GetLucky` is called again.

   2. This continues until the `index` goes out of bounds. If that happens, the `Dice` function is called with the `iteration` value incremented by 1.

4. The final number (`luckyNumber`) is calculated by using the formula `math.Mod(float64(lucky), math.Pow(10, 4))` which
   divides the value of lucky by 10^4 (10000) and gets the remainder. This ensures that the final number is in the range
   of `[0, 9999]`.

5. The function returns the random value `luckyNumber`.

#### Explanation

### Roll

#### Arguments

```go
var serverSeed string
var publicSeed string
var nonce int
// maximum represents the maximum value that can be rolled, counting from 0.
//
// Example if maximum is 5:
// 0, 1, 2, 3, 4
var maximum int
```

#### Example

```go
// Use something like a map to map the value to colors, bait, etc.
value, err := vero.Roll(serverSeed, publicSeed, nonce, maximum)
if err != nil {
  log.Fatal(err)
}
```

#### Explanation

1. The function calculates an HMAC-SHA256 hash using the `serverSeed` and a combined `seed`. This hash is used as a
   source of randomness.

2. The `GetRandomInt` function extracts a substring of length 13 from the `hash` string starting at the position `0`.

   1. This substring is then converted from a hexadecimal string to an integer.

   2. The value `e` (2^52) is calculated, which is approximately 4.5035e+15. This value represents the maximum value
      that can be represented precisely in the mantissa of a 64-bit floating-point number.

   3. The formula `math.Floor((float64(valueFromHash) / e) * float64(max))` calculates a random number that is in the
      range of `[0, max]`.

   4. The function returns the random integer and any error that may have occurred during the conversion.

3. The function returns the random value `rollValue`.

### Plinko

#### Arguments

```go
var serverSeed string
var clientSeed string
var nonce int
var iteration int
// rows represents the number of rows in the triangle.
var rows int
```

#### Example

```go
// column represents the index of the column that the ball dropped into.
column, err := vero.Plinko(serverSeed, clientSeed, nonce, 0, rows)
if err != nil {
  log.Fatal(err)
}
```

#### Explanation

```
Count it like this.

0      0
1     0 1
2    0 1 2
3   0 1 2 3
```

1. Variable `coordinate` is initialised to track the net deviation from the center position.

2. The for loop loops for any number in the range `[0, rows]`.

   1. The function calculates an HMAC-SHA256 hash using the `serverSeed` and a combined `seed`. This hash is used as a
      source of randomness.

   2. The `GetLucky` function extracts a substring of length 5 from the `hash` string starting at the
      position `index*5`. This substring is then converted from a hexadecimal string to an integer. The function
      returns the random integer and any error that may have occurred during the conversion.

   3. The for loop ensures that the `lucky` integer is above 10^6 (1000000) since it will be divided by 10^4 (10000)
      later. If it is under 10^6, then the `index` is incremented by 1 and the `GetLucky` is called again. This
      continues until the `index` goes out of bounds. If that happens, the `Dice` function is called with
      the `iteration` value incremented by 1.

   4. The final number (`luckyNumber`) is calculated by using the formula `math.Mod(float64(lucky), math.Pow(10, 4))`
      which divides the value of lucky by 10^4 (10000) and gets the remainder. This ensures that the final number is in
      the range of `[0, 9999]`.

   5. If the luckyNumber is in the range of `[0, 4999]` the ball goes to the left. (coordinate -= 1)

      If the luckyNumber is in the range of `[5000, 9999]` the ball goes to the right. (coordinate += 1)

3. The function returns the column number `(rows + coordinate) / 2` that the ball landed on.

## Documentation

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here:
https://pkg.go.dev/github.com/pastc/vero
