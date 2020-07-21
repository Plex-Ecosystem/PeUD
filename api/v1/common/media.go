package common

import "time"

type Ratings struct {
	Votes int `json:"votes"`
	Value int `json:"value"`
}
type MediaInfo struct {
	AudioBitrate     int     `json:"audioBitrate"`
	AudioChannels    float64 `json:"audioChannels"`
	AudioCodec       string  `json:"audioCodec"`
	AudioLanguages   string  `json:"audioLanguages"`
	AudioStreamCount int     `json:"audioStreamCount"`
	VideoBitDepth    int     `json:"videoBitDepth"`
	VideoBitrate     int     `json:"videoBitrate"`
	VideoCodec       string  `json:"videoCodec"`
	VideoFps         float64 `json:"videoFps"`
	Resolution       string  `json:"resolution"`
	RunTime          string  `json:"runTime"`
	ScanType         string  `json:"scanType"`
	Subtitles        string  `json:"subtitles"`
}
type QualityDefinition struct {
	Quality  Quality  `json:"quality"`
	Revision Revision `json:"revision"`
}
type Quality struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Source     string `json:"source"`
	Resolution int    `json:"resolution"`
}
type Revision struct {
	Version  int  `json:"version"`
	Real     int  `json:"real"`
	IsRepack bool `json:"isRepack"`
}
type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}
type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseWrapper struct {
	Page          int    `json:"page"`
	PageSize      int    `json:"pageSize"`
	SortKey       string `json:"sortKey"`
	SortDirection string `json:"sortDirection"`
	TotalRecords  int    `json:"totalRecords"`
}

type Action struct {
	DroppedPath        string `json:"droppedPath"`
	ImportedPath       string `json:"importedPath"`
	DownloadClient     string `json:"downloadClient"`
	DownloadClientName string `json:"downloadClientName"`
}

type Entry struct {
	QualityCutoffNotMet bool      `json:"qualityCutoffNotMet"`
	Date                time.Time `json:"date"`
	DownloadID          string    `json:"downloadId"`
	EventType           string    `json:"eventType"`
	Action              Action    `json:"data"`
	ID                  int       `json:"id"`
}

type QueueItem struct {
	Size                    float64       `json:"size"`
	Title                   string        `json:"title"`
	Sizeleft                float64       `json:"sizeleft"`
	Timeleft                string        `json:"timeleft"`
	EstimatedCompletionTime time.Time     `json:"estimatedCompletionTime"`
	Status                  string        `json:"status"`
	TrackedDownloadStatus   string        `json:"trackedDownloadStatus"`
	TrackedDownloadState    string        `json:"trackedDownloadState"`
	StatusMessages          []interface{} `json:"statusMessages"`
	DownloadID              string        `json:"downloadId"`
	Protocol                string        `json:"protocol"`
	DownloadClient          string        `json:"downloadClient"`
	Indexer                 string        `json:"indexer"`
	OutputPath              string        `json:"outputPath"`
	ID                      int           `json:"id"`
}
