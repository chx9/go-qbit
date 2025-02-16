# go-qbit
go-qbit is a lightweight and easy-to-use Go client library for interacting with the [qBittorrent Web API](https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)). It providing a simple interface to manage torrents, preferences, and more.
# Features
- Full API Coverage: Supports all major features of the qBittorrent Web API.
- Simple and Intuitive: Clean API design for easy integration into your projects.

# Installation
To install go-qbit, use the following command:

```bash
go get github.com/chx9/go-qbit
```
# QuickStart
```go
package main

import (
    "fmt"
    "log"

    "github.com/chx9/go-qbit/qbit"
)

func main() {
    // Create a new client
    client := qbit.NewClient("http://localhost:8080")

    // Log in to qBittorrent
    err := client.Login("admin", "adminadmin")
    if err != nil {
        log.Fatal(err)
    }

    opts := map[string]string{"category": "books"}
    // List torrents with optional filters, if opt is nil, list all torrents
    torrents, err := client.List(opts)
    if err != nil {
        log.Fatal(err)
    }

    // Print torrent hashes
    for _, torrent := range torrents {
        fmt.Println(torrent.Hash)
    }

}
```