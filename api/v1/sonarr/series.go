package sonarr

import (
	"time"

	"github.com/DirtyCajunRice/PeUD/api/v1/common"
)

type Series struct {
	Title             string           `json:"title"`
	AlternateTitles   []AlternateTitle `json:"alternateTitles"`
	SortTitle         string           `json:"sortTitle"`
	Status            string           `json:"status"`
	Ended             bool             `json:"ended"`
	Overview          string           `json:"overview"`
	PreviousAiring    time.Time        `json:"previousAiring"`
	Network           string           `json:"network"`
	AirTime           string           `json:"airTime"`
	Images            []Image          `json:"images"`
	Seasons           []Season         `json:"seasons"`
	Year              int              `json:"year"`
	Path              string           `json:"path"`
	QualityProfileID  int              `json:"qualityProfileId"`
	LanguageProfileID int              `json:"languageProfileId"`
	SeasonFolder      bool             `json:"seasonFolder"`
	Monitored         bool             `json:"monitored"`
	UseSceneNumbering bool             `json:"useSceneNumbering"`
	Runtime           int              `json:"runtime"`
	TvdbID            int              `json:"tvdbId"`
	TvRageID          int              `json:"tvRageId"`
	TvMazeID          int              `json:"tvMazeId"`
	FirstAired        time.Time        `json:"firstAired"`
	SeriesType        string           `json:"seriesType"`
	CleanTitle        string           `json:"cleanTitle"`
	ImdbID            string           `json:"imdbId"`
	TitleSlug         string           `json:"titleSlug"`
	RootFolderPath    string           `json:"rootFolderPath"`
	Certification     string           `json:"certification"`
	Genres            []string         `json:"genres"`
	Tags              []int            `json:"tags"`
	Added             time.Time        `json:"added"`
	Ratings           common.Ratings   `json:"ratings"`
	Statistics        SeriesStatistics `json:"statistics"`
	ID                int              `json:"id"`
}
type AlternateTitle struct {
	Title        string `json:"title"`
	SeasonNumber int    `json:"seasonNumber"`
}
type Image struct {
	common.Image
	RemoteURL string `json:"remoteUrl"`
}
type Season struct {
	SeasonNumber int              `json:"seasonNumber"`
	Monitored    bool             `json:"monitored"`
	Statistics   SeasonStatistics `json:"statistics,omitempty"`
}
type SeasonStatistics struct {
	PreviousAiring    time.Time `json:"previousAiring"`
	EpisodeFileCount  int       `json:"episodeFileCount"`
	EpisodeCount      int       `json:"episodeCount"`
	TotalEpisodeCount int       `json:"totalEpisodeCount"`
	SizeOnDisk        int64     `json:"sizeOnDisk"`
	PercentOfEpisodes int       `json:"percentOfEpisodes"`
}

type SeriesStatistics struct {
	SeasonCount       int   `json:"seasonCount"`
	EpisodeFileCount  int   `json:"episodeFileCount"`
	EpisodeCount      int   `json:"episodeCount"`
	TotalEpisodeCount int   `json:"totalEpisodeCount"`
	SizeOnDisk        int64 `json:"sizeOnDisk"`
	PercentOfEpisodes int   `json:"percentOfEpisodes"`
}
