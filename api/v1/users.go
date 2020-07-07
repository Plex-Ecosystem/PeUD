package v1

type PlexUser struct {
	ID          int    `json:"id" peud:"u,p"`
	Uuid        string `json:"uuid" peud:"u"`
	HasPassword bool   `json:"hasPassword"`
	Username    string `json:"username" peud:"u"`
	Email       string `json:"email" peud:"u"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Home        bool   `json:"home"`
	Restricted  bool   `json:"restricted"`
	Status      string `json:"status"`
	Admin       bool   `json:"admin"`
	Guest       bool   `json:"guest"`
	Protected   bool   `json:"protected"`
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
