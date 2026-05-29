package main

import (
	"github.com/sanikum3/ORP/cmd"
	"github.com/sanikum3/ORP/internal/storage"
)

func main() {
	storage.Load()
	cmd.Execute()
}
