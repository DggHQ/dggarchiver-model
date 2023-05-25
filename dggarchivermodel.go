package dggarchivermodel

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
	Platform      string `json:"platform"`
	Downloader    string `json:"downloader"`
	ID            string `json:"id"`
	PlaybackURL   string `json:"playbackurl"`
	PubTime       string `json:"pubtime"`
	Title         string `json:"title"`
	StartTime     string `json:"starttime"`
	EndTime       string `json:"endtime"`
	Thumbnail     string `json:"thumbnail"`
	ThumbnailPath string `json:"thumbnailpath"`
	Path          string `json:"path"`
	Duration      int    `json:"duration"`
}

// Contains the data structure to add the VOD into the SQLite DB
type UploadedVOD struct {
	VOD
	Claim              string
	LBRYName           string
	LBRYChannel        string
	LBRYNormalizedName string
	LBRYPermanentURL   string
}
