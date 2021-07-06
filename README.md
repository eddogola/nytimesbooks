# nytimesbooks

Wrapper around the [New York Times Books API](https://developer.nytimes.com/docs/books-product/1/overview).

[![Go](https://github.com/eddogola/nytimesbooks/actions/workflows/go.yml/badge.svg)](https://github.com/eddogola/nytimesbooks/actions/workflows/go.yml)
[![Build Status](https://travis-ci.com/eddogola/nytimesbooks.svg?branch=main)](https://travis-ci.com/eddogola/nytimesbooks)
[![codecov](https://codecov.io/gh/eddogola/nytimesbooks/branch/master/graph/badge.svg?token=dvbOf98a2V)](https://codecov.io/gh/eddogola/nytimesbooks)

## Installing

```bash
$ go get github.com/eddogola/nytimesbooks

go: finding github.com/eddogola/nytimesbooks latest
go: downloading github.com/eddogola/nytimesbooks ...
go: extracting github.com/eddogola/nytimesbooks ...
```

## Use

For example, getting the list of best sellers in the default list category(hardcover-fiction).

```go
package main

import (
    "fmt"

    "github.com/eddogola/nytimesbooks"
)

// Initialize Client
c := NewClient("apiKey")

// Make request
list, err := c.GetBestSellersList(nil)
if err != nil {
    // handle error
}

fmt.Println(list.Results)
```
