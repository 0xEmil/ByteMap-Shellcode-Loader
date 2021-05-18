package main

import (
	"github.com/0xEmil/LBD/pkg"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Encode struct {
		Key     string `arg name:"key" help:"path to encryption key"`
		Payload string `arg name:"payload" help:"path to payload"`
	} `cmd help:"Encode payload with local file."`

	Decode struct {
		Path           string `arg name:"path" help:"path of decode key file"`
		EncodedPayload string `arg name:"encodedpayload" help:"path to encoded ayload"`
	} `cmd help:"Decode a encrypted payload with local file"`
}

func main() {
	ctx := kong.Parse(&CLI)
	LocalByteDropper := pkg.LBD{}
	switch ctx.Command() {
	case "encode <key> <payload>":
		LocalByteDropper.EncodePayload(CLI.Encode.Key, CLI.Encode.Payload)
	case "decode <path> <encodedpayload>":
		LocalByteDropper.DecodePayload(CLI.Decode.Path, CLI.Decode.EncodedPayload)
	}

}
