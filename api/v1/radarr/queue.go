package radarr

import "time"

type QueueResponse struct {
	Page          int         `json:"page"`
	PageSize      int         `json:"pageSize"`
	SortKey       string      `json:"sortKey"`
	SortDirection string      `json:"sortDirection"`
	TotalRecords  int         `json:"totalRecords"`
	Queue         []QueueItem `json:"records"`
}
type QueueItem struct {
	MovieID                 int               `json:"movieId"`
	Languages               []Language        `json:"languages"`
	QualityDefinition       QualityDefinition `json:"quality"`
	CustomFormats           []CustomFormat    `json:"customFormats"`
	Size                    float64           `json:"size"`
	Title                   string            `json:"title"`
	Sizeleft                float64           `json:"sizeleft"`
	Timeleft                string            `json:"timeleft"`
	EstimatedCompletionTime time.Time         `json:"estimatedCompletionTime"`
	Status                  string            `json:"status"`
	TrackedDownloadStatus   string            `json:"trackedDownloadStatus"`
	TrackedDownloadState    string            `json:"trackedDownloadState"`
	StatusMessages          []interface{}     `json:"statusMessages"`
	DownloadID              string            `json:"downloadId"`
	Protocol                string            `json:"protocol"`
	DownloadClient          string            `json:"downloadClient"`
	Indexer                 string            `json:"indexer"`
	OutputPath              string            `json:"outputPath"`
	ID                      int               `json:"id"`
}
