package radarr

import (
	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type QueueResponse struct {
	common.ResponseWrapper
	Queue []QueueItem `json:"records"`
}
type QueueItem struct {
	MovieID           int               `json:"movieId"`
	Languages         []common.Language `json:"languages"`
	QualityDefinition QualityDefinition `json:"quality"`
	CustomFormats     []CustomFormat    `json:"customFormats"`
	common.QueueItem
}
