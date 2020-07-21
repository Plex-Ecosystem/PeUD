package tautulli

type UserResponse struct {
	BaseResponse
	Data User `json:"data"`
}
type UsersResponse struct {
	BaseResponse
	Data []User `json:"data"`
}
type User struct {
	RowID           int    `json:"row_id"`
	UserID          int    `json:"user_id"`
	Username        string `json:"username"`
	FriendlyName    string `json:"friendly_name"`
	Thumb           string `json:"thumb"`
	Email           string `json:"email"`
	IsActive        int    `json:"is_active"`
	IsAdmin         int    `json:"is_admin"`
	IsHomeUser      int    `json:"is_home_user"`
	IsAllowSync     int    `json:"is_allow_sync"`
	IsRestricted    int    `json:"is_restricted"`
	DoNotify        int    `json:"do_notify"`
	KeepHistory     int    `json:"keep_history"`
	AllowGuest      int    `json:"allow_guest"`
	ServerToken     string `json:"server_token"`
	SharedLibraries string `json:"shared_libraries"`
	FilterAll       string `json:"filter_all"`
	FilterMovies    string `json:"filter_movies"`
	FilterTv        string `json:"filter_tv"`
	FilterMusic     string `json:"filter_music"`
	FilterPhotos    string `json:"filter_photos"`
}
