package v1

type PlexUser struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	Thumb              string `json:"thumb"`
	Home               bool   `json:"home"`
	AllowTuners        bool   `json:"allowTuners"`
	AllowSync          bool   `json:"allowSync"`
	AllowCameraUpload  bool   `json:"allowCameraUpload"`
	AllowChannels      bool   `json:"allowChannels"`
	AllowSubtitleAdmin bool   `json:"allowSubtitleAdmin"`
}
