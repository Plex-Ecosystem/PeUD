package radarr

import "time"

type HistoryResponse struct {
	Page          int       `json:"page"`
	PageSize      int       `json:"pageSize"`
	SortKey       string    `json:"sortKey"`
	SortDirection string    `json:"sortDirection"`
	TotalRecords  int       `json:"totalRecords"`
	History       []History `json:"records"`
}
type History struct {
	MovieID             int               `json:"movieId"`
	SourceTitle         string            `json:"sourceTitle"`
	Languages           []Language        `json:"languages"`
	QualityDefinition   QualityDefinition `json:"quality"`
	CustomFormats       []CustomFormat    `json:"customFormats"`
	QualityCutoffNotMet bool              `json:"qualityCutoffNotMet"`
	Date                time.Time         `json:"date"`
	DownloadID          string            `json:"downloadId,omitempty"`
	EventType           string            `json:"eventType"`
	Action              Action            `json:"data,omitempty"`
	ID                  int               `json:"id"`
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
	DroppedPath        string `json:"droppedPath"`
	ImportedPath       string `json:"importedPath"`
	DownloadClient     string `json:"downloadClient"`
	DownloadClientName string `json:"downloadClientName"`
	Reason             string `json:"reason"`
}
