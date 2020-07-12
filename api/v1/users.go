package v1

type Settings struct {
	TautulliAccess string `json:"tautulliAccess"`
}

func NewSettings() *Settings {
	return &Settings{
		TautulliAccess: "global",
	}
}

type User struct {
	ID              int    `json:"id" peud:"u,p,a"`
	Alias           string `json:"alias"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	SharedLibraries string `json:""`
}

type PlexUser struct {
	ID                        int                 `json:"id" xml:"id,attr" peud:"u,p"`
	Title                     string              `json:"title" xml:"title,attr"`
	Username                  string              `json:"username" xml:"username,attr" peud:"u"`
	Email                     string              `json:"email" xml:"email,attr" peud:"u"`
	RecommendationsPlaylistId string              `json:"recommendationsPlaylistId" xml:"recommendationsPlaylistId,attr"`
	Thumb                     string              `json:"thumb" xml:"thumb,attr"`
	Protected                 bool                `json:"protected" xml:"protected,attr"`
	Home                      bool                `json:"home" xml:"home,attr"`
	AllowTuners               bool                `json:"allowTuners" xml:"allowTuners,attr"`
	AllowSync                 bool                `json:"allowSync" xml:"allowSync,attr"`
	AllowCameraUpload         bool                `json:"allowCameraUpload" xml:"allowCameraUpload,attr"`
	AllowChannels             bool                `json:"allowChannels" xml:"allowChannels,attr"`
	AllowSubtitleAdmin        bool                `json:"allowSubtitleAdmin" xml:"allowSubtitleAdmin,attr"`
	FilterAll                 string              `json:"filterAll" xml:"filterAll,attr"`
	FilterMovies              string              `json:"filterMovies" xml:"filterMovies,attr"`
	FilterMusic               string              `json:"filterMusic" xml:"filterMusic,attr"`
	FilterPhotos              string              `json:"filterPhotos" xml:"filterPhotos,attr"`
	FilterTelevision          string              `json:"filterTelevision" xml:"filterTelevision,attr"`
	Restricted                bool                `json:"restricted" xml:"restricted,attr"`
	PlexUserServers           PlexUserServerSlice `xml:"Server"`
}

type PlexUserServer struct {
	ID                int    `json:"id" xml:"id,attr"`
	ServerID          int    `json:"serverId" xml:"serverId,attr"`
	MachineIdentifier string `json:"machineIdentifier" xml:"machineIdentifier,attr"`
	Name              string `json:"name" xml:"name,attr"`
	LastSeenAt        *Time  `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	NumLibraries      int    `json:"numLibraries" xml:"numLibraries,attr"`
	AllLibraries      bool   `json:"allLibraries" xml:"allLibraries,attr"`
	Owned             bool   `json:"owned" xml:"owned,attr"`
	Pending           bool   `json:"pending" xml:"pending,attr"`
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
