package main

import (
	"flag"
	"github.com/kogonia/xdrop"
	"github.com/kogonia/xlog"
)

var (
	dropboxPath = flag.String("d", "/", "destination dropbox directory")
	filePath    = flag.String("f", "", "file to upload")
	token       = flag.String("t", "", "dropbox app token")

	help = flag.Bool("h", false, "display help message")
)

func main() {
	flag.Parse()

	if *help || len(*filePath) == 0 || len(*token) == 0 {
		flag.Usage()
		return
	}

	d, err := xdrop.New(*dropboxPath, *filePath, *token)
	if err != nil {
		xlog.Fatal(err)
	}

	err = d.Upload()
	if err != nil {
		xlog.Fatal(err)
	}
}
