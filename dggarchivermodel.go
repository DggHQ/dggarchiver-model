package dggarchivermodel

// Contains the data structure for the LiveNotify message
type LiveNotify struct {
	Live bool `json:"live"`
}

// Contains the data structure for the LiveNotifyReply message
type LiveNotifyReply struct {
	Mute bool `json:"mute"`
}

// Contains the data structure for the YouTube VOD/livestream
type YTVod struct {
	ID            string `json:"id"`
	PubTime       string `json:"pubtime"`
	Title         string `json:"title"`
	StartTime     string `json:"starttime"`
	EndTime       string `json:"endtime"`
	Thumbnail     string `json:"thumbnail"`
	ThumbnailPath string `json:"thumbnailpath"`
	Path          string `json:"path"`
	Duration      int    `json:"duration"`
}
