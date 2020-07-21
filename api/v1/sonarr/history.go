package sonarr

import (
	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type HistoryResponse struct {
	common.ResponseWrapper
	History []Entry `json:"records"`
}
type Entry struct {
	EpisodeID            int               `json:"episodeId"`
	SeriesID             int               `json:"seriesId"`
	SourceTitle          string            `json:"sourceTitle"`
	Language             common.Language   `json:"language"`
	Quality              QualityDefinition `json:"quality"`
	LanguageCutoffNotMet bool              `json:"languageCutoffNotMet"`
	common.Entry
}
