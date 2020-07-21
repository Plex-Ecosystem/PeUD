package tautulli

import "github.com/DirtyCajunRice/go-utility/types"

type UserResponse struct {
	BaseResponse
	Data User `json:"data"`
}
type UsersResponse struct {
	BaseResponse
	Data []User `json:"data"`
}
type User struct {
	RowID           int                      `json:"row_id"`
	UserID          int                      `json:"user_id"`
	Username        string                   `json:"username"`
	FriendlyName    string                   `json:"friendly_name"`
	Thumb           string                   `json:"thumb"`
	Email           string                   `json:"email"`
	IsActive        types.ConvertibleBoolean `json:"is_active"`
	IsAdmin         types.ConvertibleBoolean `json:"is_admin"`
	IsHomeUser      types.ConvertibleBoolean `json:"is_home_user"`
	IsAllowSync     types.ConvertibleBoolean `json:"is_allow_sync"`
	IsRestricted    types.ConvertibleBoolean `json:"is_restricted"`
	DoNotify        types.ConvertibleBoolean `json:"do_notify"`
	KeepHistory     types.ConvertibleBoolean `json:"keep_history"`
	AllowGuest      types.ConvertibleBoolean `json:"allow_guest"`
	ServerToken     string                   `json:"server_token"`
	SharedLibraries string                   `json:"shared_libraries"`
	FilterAll       string                   `json:"filter_all"`
	FilterMovies    string                   `json:"filter_movies"`
	FilterTv        string                   `json:"filter_tv"`
	FilterMusic     string                   `json:"filter_music"`
	FilterPhotos    string                   `json:"filter_photos"`
}
