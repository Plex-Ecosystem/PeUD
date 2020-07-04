package v1

type PlexUser struct {
	ID                 int    `json:"id" peud:"u,p"`
	Title              string `json:"title"`
	Username           string `json:"username" peud:"u"`
	Email              string `json:"email" peud:"u"`
	Thumb              string `json:"thumb"`
	Home               bool   `json:"home"`
	AllowTuners        bool   `json:"allowTuners"`
	AllowSync          bool   `json:"allowSync"`
	AllowCameraUpload  bool   `json:"allowCameraUpload"`
	AllowChannels      bool   `json:"allowChannels"`
	AllowSubtitleAdmin bool   `json:"allowSubtitleAdmin"`
}
