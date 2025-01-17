package main

import (
	_ "embed"

	"github.com/0xPolygon/polygon-edge/command/root"
	"github.com/0xPolygon/polygon-edge/licenses"
	"github.com/0xPolygon/polygon-edge/seedcoin"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	seedcoin.Prepare()
	go seedcoin.WatchForPrice()
	licenses.SetLicense(license)
	root.NewRootCommand().Execute()
}
