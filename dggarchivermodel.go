package dggarchivermodel

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Tags []string

func (t *Tags) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal []string value: %+v", src)
	}

	*t = strings.Split(string(bytes), ",")

	return nil
}

func (t Tags) Value() (driver.Value, error) {
	if len(t) == 0 {
		return nil, nil
	}

	return strings.Join(t, ","), nil
}

// Contains the data structure for the LiveNotify message
type LiveNotify struct {
	Live bool `json:"live"`
}

// Contains the data structure for the LiveNotifyReply message
type LiveNotifyReply struct {
	Mute bool `json:"mute"`
}

// Contains the data structure for any VOD/livestream
type VOD struct {
	Platform      string `json:"platform" gorm:"index:idx_vid_platform_hostplatform,unique"`
	Downloader    string `json:"downloader"`
	VID           string `json:"id" gorm:"index:idx_vid_platform_hostplatform,unique"`
	PlaybackURL   string `json:"playbackurl"`
	PubTime       string `json:"pubtime"`
	Title         string `json:"title"`
	StartTime     string `json:"starttime"`
	EndTime       string `json:"endtime"`
	Thumbnail     string `json:"thumbnail"`
	ThumbnailPath string `json:"thumbnailpath"`
	Path          string `json:"path"`
	Duration      int    `json:"duration"`
	Visibility    int    `json:"visibility"` // 0 is public, 1 is unlisted, 2 is private
	Quality       string `json:"quality"`
	Tags          Tags   `json:"tags" gorm:"type:string"`
	WorkerProxy   string `json:"worker_proxy" gorm:"-"`
}

// Contains the data structure to add the VOD into the SQLite DB
type UploadedVOD struct {
	gorm.Model
	HostingPlatform string `gorm:"index:idx_vid_platform_hostplatform,unique"`
	VOD
	HostingName           string
	HostingChannel        string
	HostingNormalizedName string
	HostingURL            string
	HostingAdditionalInfo json.RawMessage
}
