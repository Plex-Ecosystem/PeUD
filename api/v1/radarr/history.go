package radarr

import (
	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type HistoryResponse struct {
	common.ResponseWrapper
	History []Entry `json:"records"`
}
type Entry struct {
	MovieID           int               `json:"movieId"`
	SourceTitle       string            `json:"sourceTitle"`
	Languages         []common.Language `json:"languages"`
	QualityDefinition QualityDefinition `json:"quality"`
	CustomFormats     []CustomFormat    `json:"customFormats"`
	common.Entry
}
type CustomFormat struct {
	ID                              int             `json:"id"`
	Name                            string          `json:"name"`
	IncludeCustomFormatWhenRenaming bool            `json:"includeCustomFormatWhenRenaming"`
	Specifications                  []Specification `json:"specifications"`
}
type Specification struct {
	Name               string  `json:"name"`
	Implementation     string  `json:"implementation"`
	ImplementationName string  `json:"implementationName"`
	InfoLink           string  `json:"infoLink"`
	Negate             bool    `json:"negate"`
	Required           bool    `json:"required"`
	Fields             []Field `json:"fields"`
}
type Field struct {
	Order         int            `json:"order"`
	Name          string         `json:"name"`
	Label         string         `json:"label"`
	Value         int            `json:"value"`
	Type          string         `json:"type"`
	Advanced      bool           `json:"advanced"`
	SelectOptions []SelectOption `json:"selectOptions"`
}
type SelectOption struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type Action struct {
	common.Action
	Reason string `json:"reason"`
}
