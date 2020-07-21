package radarr

import "time"

type CalenderEntry struct {
	Title                 string        `json:"title"`
	AlternateTitles       []string      `json:"alternateTitles"`
	SecondaryYearSourceID int           `json:"secondaryYearSourceId"`
	SortTitle             string        `json:"sortTitle"`
	SizeOnDisk            int           `json:"sizeOnDisk"`
	Status                string        `json:"status"`
	Overview              string        `json:"overview"`
	InCinemas             time.Time     `json:"inCinemas"`
	DigitalRelease        time.Time     `json:"digitalRelease"`
	Images                []Image       `json:"images"`
	Website               string        `json:"website"`
	Year                  int           `json:"year"`
	HasFile               bool          `json:"hasFile"`
	YouTubeTrailerID      string        `json:"youTubeTrailerId"`
	Studio                string        `json:"studio"`
	Path                  string        `json:"path"`
	QualityProfileID      int           `json:"qualityProfileId"`
	Monitored             bool          `json:"monitored"`
	MinimumAvailability   string        `json:"minimumAvailability"`
	IsAvailable           bool          `json:"isAvailable"`
	FolderName            string        `json:"folderName"`
	Runtime               int           `json:"runtime"`
	CleanTitle            string        `json:"cleanTitle"`
	ImdbID                string        `json:"imdbId"`
	TmdbID                int           `json:"tmdbId"`
	TitleSlug             string        `json:"titleSlug"`
	Genres                []string      `json:"genres"`
	Tags                  []interface{} `json:"tags"`
	Added                 time.Time     `json:"added"`
	Ratings               Ratings       `json:"ratings"`
	ID                    int           `json:"id"`
}
