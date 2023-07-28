package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Almazatun/qrcode-iod/pkg/command"
	"github.com/Almazatun/qrcode-iod/pkg/common"
)

func main() {
	//cli cmd
	generateCMD := flag.NewFlagSet("gen", flag.ExitOnError)
	addWebURL := generateCMD.String("url", "", "generate QR Code by website url")

	if len(os.Args) < 2 {
		fmt.Println("🤖 Expected 'gen' subcommands")
		os.Exit(1)
	}

	// cmds
	qr := command.GenQECode{}

	switch os.Args[1] {
	case "gen":
		qr.Create(generateCMD, addWebURL)
	default:
		fmt.Println("Not found command")
		fmt.Println(common.COW_SHAPE)
	}
}
