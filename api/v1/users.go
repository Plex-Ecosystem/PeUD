package v1

type User struct {
	ID              int    `json:"id" peud:"u,p"`
	Alias           string `json:"alias"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	SharedLibraries string `json:""`
}

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
	RowID            int                `json:"row_id" peud:"u,p"`
	Username         string             `json:"username" peud:"u"`
	Email            string             `json:"email" peud:"u"`
	Thumb            string             `json:"thumb"`
	FilterAll        string             `json:"filter_all"`
	FilterMovies     string             `json:"filter_movies"`
	FilterMusic      string             `json:"filter_music"`
	FilterPhotos     string             `json:"filter_photos"`
	FilterTelevision string             `json:"filter_television"`
	UserID           int                `json:"user_id" peud:"u,p"`
	FriendlyName     string             `json:"friendly_name"`
	IsActive         ConvertibleBoolean `json:"is_active"`
	IsAdmin          ConvertibleBoolean `json:"is_admin"`
	IsHomeUser       ConvertibleBoolean `json:"is_home_user"`
	IsAllowSync      ConvertibleBoolean `json:"is_allow_sync"`
	IsRestricted     ConvertibleBoolean `json:"is_restricted"`
	DoNotify         ConvertibleBoolean `json:"do_notify"`
	KeepHistory      ConvertibleBoolean `json:"keep_history"`
	AllowGuest       ConvertibleBoolean `json:"allow_guest"`
	ServerToken      string             `json:"server_token"`
	SharedLibraries  string             `json:"shared_libraries"`
}

type OrganizrUser struct {
	ID           int    `json:"id" peud:"u,p"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PlexToken    string `json:"plex_token"`
	Group        string `json:"group"`
	GroupID      int    `json:"group_id"`
	Locked       string `json:"locked,omitempty"`
	Image        string `json:"image"`
	RegisterDate string `json:"register_date"`
	AuthService  string `json:"auth_service"`
}

type OmbiUser struct {
	ID                        string `json:"id" peud:"u,p"`
	Username                  string `json:"userName"`
	Alias                     string `json:"alias"`
	EmailAddress              string `json:"emailAddress"`
	Password                  string `json:"password"`
	LastLoggedIn              string `json:"lastLoggedIn"`
	Language                  string `json:"language"`
	HasLoggedIn               bool   `json:"hasLoggedIn"`
	UserType                  int    `json:"userType"`
	MovieRequestLimit         int    `json:"movieRequestLimit"`
	EpisodeRequestLimit       int    `json:"episodeRequestLimit"`
	EpisodeRequestQuota       string `json:"episodeRequestQuota"`
	MovieRequestQuota         string `json:"movieRequestQuota"`
	MusicRequestQuota         string `json:"musicRequestQuota"`
	MusicRequestLimit         int    `json:"musicRequestLimit"`
	RequestTv                 bool   `json:"requestTV"`
	RequestMovie              bool   `json:"requestMovie"`
	AutoApproveMovie          bool   `json:"autoApproveMovie"`
	Admin                     bool   `json:"admin"`
	AutoApproveTv             bool   `json:"autoApproveTv"`
	AutoApproveMusic          bool   `json:"autoApproveMusic"`
	RequestMusic              bool   `json:"requestMusic"`
	PowerUser                 bool   `json:"powerUser"`
	Disabled                  bool   `json:"disabled"`
	ReceivesNewsletter        bool   `json:"receivesNewsletter"`
	ManageOwnRequests         bool   `json:"manageOwnRequests"`
	EditCustomPage            bool   `json:"editCustomPage"`
	UserID                    string `json:"userId"`
	SonarrQualityProfileAnime int    `json:"sonarrQualityProfileAnime"`
	SonarrRootPathAnime       int    `json:"sonarrRootPathAnime"`
	SonarrRootPath            int    `json:"sonarrRootPath"`
	SonarrQualityProfile      int    `json:"sonarrQualityProfile"`
	RadarrRootPath            int    `json:"radarrRootPath"`
	RadarrQualityProfile      int    `json:"radarrQualityProfile"`
}
