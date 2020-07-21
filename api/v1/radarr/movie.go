package radarr

import (
	"time"

	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type Movie struct {
	Title                 string           `json:"title"`
	AlternateTitles       []AlternateTitle `json:"alternateTitles"`
	SecondaryYearSourceID int              `json:"secondaryYearSourceId"`
	SortTitle             string           `json:"sortTitle"`
	SizeOnDisk            int64            `json:"sizeOnDisk"`
	Status                string           `json:"status"`
	Overview              string           `json:"overview"`
	InCinemas             time.Time        `json:"inCinemas"`
	PhysicalRelease       time.Time        `json:"physicalRelease"`
	Images                []common.Image   `json:"images"`
	Website               string           `json:"website"`
	Year                  int              `json:"year"`
	HasFile               bool             `json:"hasFile"`
	YouTubeTrailerID      string           `json:"youTubeTrailerId"`
	Studio                string           `json:"studio"`
	Path                  string           `json:"path"`
	QualityProfileID      int              `json:"qualityProfileId"`
	Monitored             bool             `json:"monitored"`
	MinimumAvailability   string           `json:"minimumAvailability"`
	IsAvailable           bool             `json:"isAvailable"`
	FolderName            string           `json:"folderName"`
	Runtime               int              `json:"runtime"`
	CleanTitle            string           `json:"cleanTitle"`
	ImdbID                string           `json:"imdbId"`
	TmdbID                int              `json:"tmdbId"`
	TitleSlug             string           `json:"titleSlug"`
	Certification         string           `json:"certification"`
	Genres                []string         `json:"genres"`
	Tags                  []string         `json:"tags"`
	Added                 time.Time        `json:"added"`
	Ratings               common.Ratings   `json:"ratings"`
	MovieFile             MovieFile        `json:"movieFile"`
	ID                    int              `json:"id"`
}
type AlternateTitle struct {
	SourceType string          `json:"sourceType"`
	MovieID    int             `json:"movieId"`
	Title      string          `json:"title"`
	SourceID   int             `json:"sourceId"`
	Votes      int             `json:"votes"`
	VoteCount  int             `json:"voteCount"`
	Language   common.Language `json:"language"`
	ID         int             `json:"id"`
}

type MovieFile struct {
	MovieID             int               `json:"movieId"`
	RelativePath        string            `json:"relativePath"`
	Path                string            `json:"path"`
	Size                int64             `json:"size"`
	DateAdded           time.Time         `json:"dateAdded"`
	SceneName           string            `json:"sceneName"`
	IndexerFlags        int               `json:"indexerFlags"`
	QualityDefinition   QualityDefinition `json:"quality"`
	MediaInfo           MediaInfo         `json:"mediaInfo"`
	QualityCutoffNotMet bool              `json:"qualityCutoffNotMet"`
	Languages           []common.Language `json:"languages"`
	ID                  int               `json:"id"`
}
type QualityDefinition struct {
	Quality Quality `json:"quality"`
	common.QualityDefinition
}
type Quality struct {
	common.Quality
	Modifier string `json:"modifier"`
}
type MediaInfo struct {
	AudioAdditionalFeatures string `json:"audioAdditionalFeatures"`
	common.MediaInfo
}
