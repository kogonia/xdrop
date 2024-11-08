package main

import (
	"flag"
	"github.com/kogonia/xdrop"
	"github.com/kogonia/xlog"
)

const dflDropboxPath = "/app/"
const dflFilePath = "test.fb2"
const dflToken = "token"

var dropboxPath = flag.String("d", dflDropboxPath, "destination dropbox directory")
var filePath = flag.String("f", dflFilePath, "file to upload")
var token = flag.String("t", dflToken, "dropbox app token")

func main() {
	flag.Parse()
	if len(*filePath) == 0 || len(*token) == 0 {
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
