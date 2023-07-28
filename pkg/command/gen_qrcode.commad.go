package command

import (
	"bytes"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/Almazatun/qrcode-iod/pkg/common"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type GenQECode struct{}

func (qr *GenQECode) Create(genCMD *flag.FlagSet, u *string) {
	genCMD.Parse(os.Args[2:])

	if *u == "" {
		fmt.Print("🤖 Website ULR is required \n")
		genCMD.PrintDefaults()
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(*u)

	if err != nil {
		fmt.Print("🤖 Invalid website URL \n")
		genCMD.PrintDefaults()
		os.Exit(1)
	}

	qrc, err := qrcode.New(*u)

	if err != nil {
		fmt.Print("🤖" + err.Error() + " \n")
		genCMD.PrintDefaults()
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)
	wr := common.NopCloser{Writer: buf}
	w2 := standard.NewWithWriter(wr, standard.WithQRWidth(1))

	if err = qrc.Save(w2); err != nil {
		fmt.Print("🤖" + err.Error() + " \n")
		genCMD.PrintDefaults()
		os.Exit(1)
	}
}
