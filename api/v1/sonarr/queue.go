package sonarr

import (
	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type QueueResponse struct {
	common.ResponseWrapper
	Queue []QueueItem `json:"records"`
}

type QueueItem struct {
	SeriesID  int                      `json:"seriesId"`
	EpisodeID int                      `json:"episodeId"`
	Language  common.Language          `json:"language"`
	Quality   common.QualityDefinition `json:"quality"`
	common.QueueItem
}
