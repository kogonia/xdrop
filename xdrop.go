package xdrop

import (
	"github.com/kogonia/xlog"
	"net/http"
	"os"
	"time"
)

type DropBox struct {
	ApiUrl        string
	Authorisation string
	ContentType   string
	DropboxApiArg *DropboxApiArg
	DataBinary    *os.File
}

const dflMode = "add"

var client = &http.Client{
	Transport: &http.Transport{
		DisableCompression: true},
	Timeout: time.Second * 60,
}

func New(dropboxPath, filePath, fileName, token string) (*DropBox, error) {
	file, err := binaryData(filePath)
	if err != nil {
		xlog.Errorf("fail to open file \"%s\": %v", filePath, err)
		return nil, err
	}

	return &DropBox{
		ApiUrl:        "https://content.dropboxapi.com/2/files/upload",
		Authorisation: "Bearer " + token,
		ContentType:   "application/octet-stream",
		DropboxApiArg: &DropboxApiArg{
			AutoRename:     false,
			Mode:           dflMode,
			Mute:           false,
			Path:           dropboxPath + fileName,
			StrictConflict: false,
		},
		DataBinary: file,
	}, nil

}

func (d *DropBox) Upload() error {
	req, err := http.NewRequest(http.MethodPost, d.ApiUrl, d.DataBinary)
	if err != nil {
		xlog.Errorf("fail to create request: %v", err)
		return err
	}

	req.Header.Set("Authorization", d.Authorisation)
	req.Header.Set("Content-Type", d.ContentType)
	req.Header.Set("Dropbox-API-Arg", d.DropboxApiArg.String())

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("fail to send request: %v", err)
		return err
	}
	defer resp.Body.Close()

	return handleResponse(resp.Body)
}

func binaryData(filePath string) (*os.File, error) {
	return os.Open(filePath)
}
