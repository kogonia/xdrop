package main

import (
	"flag"
	"github.com/kogonia/xdrop"
	"github.com/kogonia/xlog"
	"path/filepath"
)

var (
	dropboxPath = flag.String("d", "/", "destination dropbox directory")
	filePath    = flag.String("f", "", "file to upload")
	token       = flag.String("t", "", "dropbox app token")

	preservePath = flag.Bool("p", false, "preserve path. If not set file path will be removed from destination file name")
	help         = flag.Bool("h", false, "display help message")
)

func main() {
	flag.Parse()

	if *help || len(*filePath) == 0 || len(*token) == 0 {
		flag.Usage()
		return
	}

	var name string
	if !*preservePath {
		name = filepath.Base(*filePath)
	}
	d, err := xdrop.New(*dropboxPath, *filePath, name, *token)
	if err != nil {
		xlog.Fatal(err)
	}

	err = d.Upload()
	if err != nil {
		xlog.Fatal(err)
	}
}
