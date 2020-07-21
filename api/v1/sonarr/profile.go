package sonarr

import "github.com/DirtyCajunRice/PeUD/api/v1/common"

type Profile struct {
	Name           string `json:"name"`
	UpgradeAllowed bool   `json:"upgradeAllowed"`
	ID             int    `json:"id"`
}

type QualityProfile struct {
	Profile
	Cutoff         int             `json:"cutoff"`
	QualityConfigs []QualityConfig `json:"items"`
}
type QualityConfig struct {
	Quality common.Quality `json:"quality,omitempty"`
	Items   []interface{}  `json:"items"`
	Allowed bool           `json:"allowed"`
	Name    string         `json:"name,omitempty"`
	ID      int            `json:"id,omitempty"`
}

type LanguageProfile struct {
	Profile
	Cutoff    Cutoff               `json:"cutoff"`
	Languages []LanguageDefinition `json:"languages"`
}
type Cutoff struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type LanguageDefinition struct {
	Language common.Language `json:"language"`
	Allowed  bool            `json:"allowed"`
}

type DelayProfile struct {
	EnableUsenet      bool   `json:"enableUsenet"`
	EnableTorrent     bool   `json:"enableTorrent"`
	PreferredProtocol string `json:"preferredProtocol"`
	UsenetDelay       int    `json:"usenetDelay"`
	TorrentDelay      int    `json:"torrentDelay"`
	Order             int64  `json:"order"`
	Tags              []int  `json:"tags"`
	ID                int    `json:"id"`
}

type Tag struct {
	Label string `json:"label"`
	ID    int    `json:"id"`
}

type ReleaseProfile struct {
	Enabled                      bool            `json:"enabled"`
	Required                     string          `json:"required"`
	Ignored                      string          `json:"ignored"`
	PreferredTerms               []PreferredTerm `json:"preferred"`
	IncludePreferredWhenRenaming bool            `json:"includePreferredWhenRenaming"`
	IndexerID                    int             `json:"indexerId"`
	Tags                         []int           `json:"tags"`
	ID                           int             `json:"id"`
}
type PreferredTerm struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type Indexer []struct {
	EnableRss               bool             `json:"enableRss"`
	EnableAutomaticSearch   bool             `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool             `json:"enableInteractiveSearch"`
	SupportsRss             bool             `json:"supportsRss"`
	SupportsSearch          bool             `json:"supportsSearch"`
	Protocol                string           `json:"protocol"`
	Name                    string           `json:"name"`
	Settings                []IndexerSetting `json:"fields"`
	ImplementationName      string           `json:"implementationName"`
	Implementation          string           `json:"implementation"`
	ConfigContract          string           `json:"configContract"`
	InfoLink                string           `json:"infoLink"`
	Tags                    []int            `json:"tags"`
	ID                      int              `json:"id"`
}
type IndexerSetting struct {
	Order    int    `json:"order"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	HelpText string `json:"helpText,omitempty"`
	Value    string `json:"value,omitempty"`
	Type     string `json:"type"`
	Advanced bool   `json:"advanced"`
	Unit     string `json:"unit,omitempty"`
}
