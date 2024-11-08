package xdrop

import (
	"encoding/json"
	"github.com/kogonia/xlog"
	"io"
	"strings"
	"time"
)

type response struct {
	uploadResponse
	uploadError
}

type uploadResponse struct {
	ClientModified time.Time `json:"client_modified"`
	ContentHash    string    `json:"content_hash"`
	FileLockInfo   struct {
		Created        time.Time `json:"created"`
		IsLockholder   bool      `json:"is_lockholder"`
		LockholderName string    `json:"lockholder_name"`
	} `json:"file_lock_info"`
	HasExplicitSharedMembers bool   `json:"has_explicit_shared_members"`
	Id                       string `json:"id"`
	IsDownloadable           bool   `json:"is_downloadable"`
	Name                     string `json:"name"`
	PathDisplay              string `json:"path_display"`
	PathLower                string `json:"path_lower"`
	PropertyGroups           []struct {
		Fields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
		TemplateId string `json:"template_id"`
	} `json:"property_groups"`
	Rev            string    `json:"rev"`
	ServerModified time.Time `json:"server_modified"`
	SharingInfo    struct {
		ModifiedBy           string `json:"modified_by"`
		ParentSharedFolderId string `json:"parent_shared_folder_id"`
		ReadOnly             bool   `json:"read_only"`
	} `json:"sharing_info"`
	Size int `json:"size"`
}

type uploadError struct {
	ErrorTag struct {
		Tag string `json:".tag"`
	} `json:"error"`
	ErrorSummary string `json:"error_summary"`
}

func (e uploadError) Error() string {
	sb := strings.Builder{}
	sb.WriteString(e.ErrorTag.Tag)
	sb.WriteString(": ")
	sb.WriteString(e.ErrorSummary)

	return sb.String()
}

func handleResponse(body io.ReadCloser) error {
	var resp = &response{}

	dec := json.NewDecoder(body)
	if err := dec.Decode(resp); err != nil {
		return err
	}

	if len(resp.uploadError.ErrorTag.Tag) != 0 {
		return resp.uploadError
	}
	xlog.Printf("%s [%d bytes]", resp.uploadResponse.PathDisplay, resp.uploadResponse.Size)
	return nil
}
