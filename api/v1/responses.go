package v1

import "reflect"

type PlexResponse struct {
	PlexUsers []PlexUser `xml:"User"`
}

type TautulliResponse struct {
	Response struct {
		Data []TautulliUser `json:"data"`
	} `json:"response"`
}

type OrganizrResponse struct {
	Data struct {
		Users []OrganizrUser `json:"users"`
	} `json:"data"`
}

type OmbiUserResponse struct {
	ID                  string                   `json:"id" peud:"u,p"`
	Username            string                   `json:"userName"`
	Alias               string                   `json:"alias"`
	Claims              []map[string]interface{} `json:"claims"`
	EmailAddress        string                   `json:"emailAddress"`
	Password            string                   `json:"password"`
	LastLoggedIn        string                   `json:"lastLoggedIn"`
	Language            string                   `json:"language"`
	HasLoggedIn         bool                     `json:"hasLoggedIn"`
	UserType            int                      `json:"userType"`
	MovieRequestLimit   int                      `json:"movieRequestLimit"`
	EpisodeRequestLimit int                      `json:"episodeRequestLimit"`
	EpisodeRequestQuota string                   `json:"episodeRequestQuota"`
	MovieRequestQuota   string                   `json:"movieRequestQuota"`
	MusicRequestQuota   string                   `json:"musicRequestQuota"`
	MusicRequestLimit   int                      `json:"musicRequestLimit"`
	UserQualityProfiles struct {
		UserId                    string `json:"userId"`
		SonarrQualityProfileAnime int    `json:"sonarrQualityProfileAnime"`
		SonarrRootPathAnime       int    `json:"sonarrRootPathAnime"`
		SonarrRootPath            int    `json:"sonarrRootPath"`
		SonarrQualityProfile      int    `json:"sonarrQualityProfile"`
		RadarrRootPath            int    `json:"radarrRootPath"`
		RadarrQualityProfile      int    `json:"radarrQualityProfile"`
		ID                        int    `json:"id"`
	} `json:"userQualityProfiles"`
}

func (r *OmbiUserResponse) ConvertToSane() (u OmbiUser) {
	u = OmbiUser{
		ID:                        r.ID,
		Username:                  r.Username,
		Alias:                     r.Alias,
		EmailAddress:              r.EmailAddress,
		Password:                  r.Password,
		LastLoggedIn:              r.LastLoggedIn,
		Language:                  r.Language,
		HasLoggedIn:               r.HasLoggedIn,
		UserType:                  r.UserType,
		MovieRequestLimit:         r.MovieRequestLimit,
		EpisodeRequestLimit:       r.EpisodeRequestLimit,
		EpisodeRequestQuota:       r.EpisodeRequestQuota,
		MovieRequestQuota:         r.MovieRequestQuota,
		MusicRequestQuota:         r.MovieRequestQuota,
		MusicRequestLimit:         r.MusicRequestLimit,
		UserID:                    r.UserQualityProfiles.UserId,
		SonarrQualityProfileAnime: r.UserQualityProfiles.SonarrQualityProfileAnime,
		SonarrRootPathAnime:       r.UserQualityProfiles.SonarrRootPathAnime,
		SonarrRootPath:            r.UserQualityProfiles.SonarrRootPath,
		SonarrQualityProfile:      r.UserQualityProfiles.SonarrQualityProfile,
		RadarrRootPath:            r.UserQualityProfiles.RadarrRootPath,
		RadarrQualityProfile:      r.UserQualityProfiles.RadarrQualityProfile,
	}
	for _, k := range r.Claims {
		reflect.ValueOf(&u).Elem().FieldByName(k["value"].(string)).SetBool(k["enabled"].(bool))
	}
	return
}
