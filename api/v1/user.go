package v1

type PlexUser struct {
	Username                  string `json:"username" peud:"u"`
	Email                     string `json:"email" peud:"u"`
	Thumb                     string `json:"thumb"`
	FilterAll                 string `json:"filterAll"`
	FilterMovies              string `json:"filterMovies"`
	FilterMusic               string `json:"filterMusic"`
	FilterPhotos              string `json:"filterPhotos"`
	FilterTelevision          string `json:"filterTelevision"`
	ID                        int    `json:"id" peud:"u,p"`
	Title                     string `json:"title"`
	RecommendationsPlaylistId string `json:"recommendationsPlaylistId"`
	Home                      bool   `json:"home"`
	AllowTuners               bool   `json:"allowTuners"`
	AllowSync                 bool   `json:"allowSync"`
	AllowCameraUpload         bool   `json:"allowCameraUpload"`
	AllowChannels             bool   `json:"allowChannels"`
	AllowSubtitleAdmin        bool   `json:"allowSubtitleAdmin"`
	Restricted                bool   `json:"restricted"`
}

type TautulliUser struct {
	RowID            int    `json:"row_id" peud:"u"`
	Username         string `json:"username" peud:"u"`
	Email            string `json:"email" peud:"u"`
	Thumb            string `json:"thumb"`
	FilterAll        string `json:"filter_all"`
	FilterMovies     string `json:"filter_movies"`
	FilterMusic      string `json:"filter_music"`
	FilterPhotos     string `json:"filter_photos"`
	FilterTelevision string `json:"filter_television"`
	UserID           int    `json:"user_id" peud:"u,p"`
	FriendlyName     string `json:"friendly_name"`
	IsActive         bool   `json:"is_active"`
	IsAdmin          bool   `json:"is_admin"`
	IsHomeUser       bool   `json:"is_home_user"`
	IsAllowSync      bool   `json:"is_allow_sync"`
	IsRestricted     bool   `json:"is_restricted"`
	DoNotify         bool   `json:"do_notify"`
	KeepHistory      bool   `json:"keep_history"`
	AllowGuest       bool   `json:"allow_guest"`
	ServerToken      string `json:"server_token"`
	SharedLibraries  string `json:"shared_libraries"`
}
