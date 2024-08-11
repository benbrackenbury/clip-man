package main

import (
	"github.com/benbrackenbury/clipman/src/store"
	"github.com/benbrackenbury/clipman/src/transmit"
)

func main() {
    store := store.NewLogFileStore("clipboard.log")
    defer store.Close()
    transmit.Transmit(store)
}

