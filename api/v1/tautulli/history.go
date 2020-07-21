package tautulli

type HistoryResponse struct {
	BaseResponse
	HistoryQuery HistoryQuery `json:"data"`
}
type HistoryQuery struct {
	RecordsFiltered int       `json:"recordsFiltered"`
	RecordsTotal    int       `json:"recordsTotal"`
	History         []History `json:"data"`
	Draw            int       `json:"draw"`
	FilterDuration  string    `json:"filter_duration"`
	TotalDuration   string    `json:"total_duration"`
}
type History struct {
	ReferenceID           int    `json:"reference_id"`
	RowID                 int    `json:"row_id"`
	ID                    int    `json:"id"`
	Date                  int    `json:"date"`
	Started               int    `json:"started"`
	Stopped               int    `json:"stopped"`
	Duration              int    `json:"duration"`
	PausedCounter         int    `json:"paused_counter"`
	UserID                int    `json:"user_id"`
	User                  string `json:"user"`
	FriendlyName          string `json:"friendly_name"`
	Platform              string `json:"platform"`
	Product               string `json:"product"`
	Player                string `json:"player"`
	IPAddress             string `json:"ip_address"`
	Live                  int    `json:"live"`
	MediaType             string `json:"media_type"`
	RatingKey             int    `json:"rating_key"`
	ParentRatingKey       int    `json:"parent_rating_key"`
	GrandparentRatingKey  int    `json:"grandparent_rating_key"`
	FullTitle             string `json:"full_title"`
	Title                 string `json:"title"`
	ParentTitle           string `json:"parent_title"`
	GrandparentTitle      string `json:"grandparent_title"`
	OriginalTitle         string `json:"original_title"`
	Year                  int    `json:"year"`
	MediaIndex            int    `json:"media_index"`
	ParentMediaIndex      int    `json:"parent_media_index"`
	Thumb                 string `json:"thumb"`
	OriginallyAvailableAt string `json:"originally_available_at"`
	GUID                  string `json:"guid"`
	TranscodeDecision     string `json:"transcode_decision"`
	PercentComplete       int    `json:"percent_complete"`
	WatchedStatus         int    `json:"watched_status"`
	GroupCount            int    `json:"group_count"`
	GroupIds              string `json:"group_ids"`
	State                 string `json:"state"`
	SessionKey            int    `json:"session_key"`
}
