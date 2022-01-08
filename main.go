package main

import (
	"github.com/OpachevskyEmil/ByteMap-Shellcode-Loader/pkg"
	"github.com/alecthomas/kong"
)

var CLI struct {
	Encode struct {
		Key         string `arg name:"key" help:"path to encryption key"`
		PayloadName string `arg name:"payloadname" help:"filename of the encoded payload"`
		Payload     string `arg name:"payload" help:"shellcode payload"`
	} `cmd help:"Encode payload with local file."`

	DecodeRun struct {
		Path           string `arg name:"path" help:"path of decode key file"`
		EncodedPayload string `arg name:"encodedpayload" help:"path to encoded ayload"`
	} `cmd help:"Decode a encrypted payload with local file"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "encode <key> <payloadname> <payload>":
		pkg.EncodePayload(CLI.Encode.Key, CLI.Encode.PayloadName, CLI.Encode.Payload)
	case "decode-run <path> <encodedpayload>":
		shellcode := pkg.DecodePayload(CLI.DecodeRun.Path, CLI.DecodeRun.EncodedPayload)
		pkg.Run(shellcode.Bytes())
	}
}
