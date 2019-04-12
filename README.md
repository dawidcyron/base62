# base62
Simple base62 encoding and decoding library for Golang 

## Install
`go get -u github.com/dawidcyron/base62`

## Usage
```go
package main

import (
	"log"
	"github.com/dawidcyron/base62"
)

func main() {
  //Number to encode
  toEncode := 5201
  encoded := base62.ToBase62(toEncode)
  log.Print(encoded)
  decoded := base62.ToBase10(encoded)
  log.Print(decoded)
}
```
