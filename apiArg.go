package xdrop

import (
	"encoding/json"
	"github.com/kogonia/xlog"
)

type DropboxApiArg struct {
	AutoRename     bool   `json:"autorename"`
	Mode           string `json:"mode"`
	Mute           bool   `json:"mute"`
	Path           string `json:"path"`
	StrictConflict bool   `json:"strict_conflict"`
}

func (da *DropboxApiArg) String() string {
	b, err := json.Marshal(da)
	if err != nil {
		xlog.Errorf("failed marshal json for %v: %v", *da, err)
		return ""
	}
	return string(b)
}
