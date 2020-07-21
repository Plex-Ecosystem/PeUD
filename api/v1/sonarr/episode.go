package sonarr

import (
	"time"

	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type EpisodeBase struct {
	SeriesID     int `json:"seriesId"`
	SeasonNumber int `json:"seasonNumber"`
	ID           int `json:"id"`
}

type Episode struct {
	EpisodeBase
	EpisodeFileID            int       `json:"episodeFileId"`
	EpisodeNumber            int       `json:"episodeNumber"`
	Title                    string    `json:"title"`
	Overview                 string    `json:"overview"`
	HasFile                  bool      `json:"hasFile"`
	Monitored                bool      `json:"monitored"`
	UnverifiedSceneNumbering bool      `json:"unverifiedSceneNumbering"`
	AirDate                  string    `json:"airDate,omitempty"`
	AirDateUtc               time.Time `json:"airDateUtc,omitempty"`
	AbsoluteEpisodeNumber    int       `json:"absoluteEpisodeNumber,omitempty"`
}
type EpisodeFile struct {
	EpisodeBase
	RelativePath         string                   `json:"relativePath"`
	Path                 string                   `json:"path"`
	Size                 int64                    `json:"size"`
	DateAdded            time.Time                `json:"dateAdded"`
	SceneName            string                   `json:"sceneName"`
	Language             common.Language          `json:"language"`
	QualityDefinition    common.QualityDefinition `json:"quality"`
	MediaInfo            common.MediaInfo         `json:"mediaInfo"`
	QualityCutoffNotMet  bool                     `json:"qualityCutoffNotMet"`
	LanguageCutoffNotMet bool                     `json:"languageCutoffNotMet"`
}
