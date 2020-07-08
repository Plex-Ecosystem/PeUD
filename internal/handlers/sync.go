package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

type TautulliResponse struct {
	Response struct {
		Data []v1.TautulliUser `json:"data"`
	} `json:"response"`
}

type OrganizrResponse struct {
	Data struct {
		Users []v1.OrganizrUser `json:"users"`
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

func (r *OmbiUserResponse) convertToSane() (u v1.OmbiUser) {
	u = v1.OmbiUser{
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

func Sync(env *Env, w http.ResponseWriter, r *http.Request) {
	filter := strings.Split(r.URL.Query().Get("only"), ",")
	if len(filter) == 0 {
		filter = []string{"organizr", "tautulli", "plex", "ombi"}
	}
	for _, i := range filter {
		sync(i, env)
	}
}

func sharedRequest(c *http.Client, u string, h map[string][]string, l *logrus.Entry, e string) []byte {
	r, _ := http.NewRequest("GET", u, nil)
	if e != "tautulli" {
		r.Header = h
	}
	r.Header.Add("Accept", "application/json")
	start := time.Now()
	resp, err := c.Do(r)
	if err != nil {
		l.Error(err)
	}
	l.WithFields(logrus.Fields{
		"request": "upstream",
		"api":     e,
		"took":    time.Since(start).Nanoseconds(),
	}).Debug("api call successful")
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error(err)
	}
	return b
}

func sync(e string, env *Env) {
	l := env.Log
	c := &http.Client{}
	var (
		auth = env.Config.Authentication
		db   = env.Config.Database
		h    map[string][]string
	)
	switch e {
	case "plex":
		u := "https://plex.tv/api/v2/friends"
		h = map[string][]string{
			"X-Plex-Client-Identifier": {"PeUD"},
			"X-Plex-Token":             {auth.PlexToken},
		}
		b := sharedRequest(c, u, h, l, e)
		plexFriends := make([]v1.PlexUser, 0)
		if err := json.Unmarshal(b, &plexFriends); err != nil {
			l.Error(err)
		}
		db.InsertUsers("plexUsers", plexFriends)
	case "tautulli":
		u := fmt.Sprintf("%s/api/v2?cmd=get_users&apikey=%s", auth.TautulliURL, auth.TautulliKey)
		b := sharedRequest(c, u, h, l, e)
		tautulliResponse := TautulliResponse{}
		if err := json.Unmarshal(b, &tautulliResponse); err != nil {
			l.Error(err)
		}
		db.InsertUsers("tautulliUsers", tautulliResponse.Response.Data)
	case "organizr":
		u := fmt.Sprintf("%s/api?v1/user/list", auth.OrganizrURL)
		h = map[string][]string{"token": {auth.OrganizrToken}}
		b := sharedRequest(c, u, h, l, e)
		organizrResponse := OrganizrResponse{}
		if err := json.Unmarshal(b, &organizrResponse); err != nil {
			l.Error(err)
		}
		db.InsertUsers("organizrUsers", organizrResponse.Data.Users)
	case "ombi":
		u := fmt.Sprintf("%s/api/v1/Identity/Users", auth.OmbiURL)
		h = map[string][]string{"ApiKey": {auth.OmbiKey}}
		b := sharedRequest(c, u, h, l, e)
		ombiUserResponse := make([]OmbiUserResponse, 0)
		ombiUsers := make([]v1.OmbiUser, 0)
		if err := json.Unmarshal(b, &ombiUserResponse); err != nil {
			l.Error(err)
		}
		for _, r := range ombiUserResponse {
			ombiUsers = append(ombiUsers, r.convertToSane())
		}
		db.InsertUsers("ombiUsers", ombiUsers)
	}
}
