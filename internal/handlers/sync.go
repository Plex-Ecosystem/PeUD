package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func Sync(env *Env, w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("only")
	if strings.Contains(filter, "plex") || filter == "" {
		syncPlex(env)
	}
	if strings.Contains(filter, "tautulli") || filter == "" {
		syncTautulli(env)
	}
	if strings.Contains(filter, "organizr") || filter == "" {
		syncOrganizr(env)
	}
}

func syncPlex(env *Env) {
	log := env.Log
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://plex.tv/api/v2/friends", nil)
	req.Header.Add("X-Plex-Client-Identifier", "PeUD")
	req.Header.Add("X-Plex-Token", env.Config.Authentication.PlexToken)
	req.Header.Add("Accept", "application/json")
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	log.WithFields(logrus.Fields{
		"request": "upstream",
		"api":     "plex.tv",
		"took":    time.Since(start).Nanoseconds(),
	}).Debug("plex api call successful")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	plexFriends := make([]v1.PlexUser, 0)
	if err := json.Unmarshal(body, &plexFriends); err != nil {
		log.Error(err)
	}
	env.Config.Database.InsertPlexUsers(plexFriends)
}

func syncTautulli(env *Env) {
	log := env.Log
	client := &http.Client{}
	auth := env.Config.Authentication
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v2?cmd=get_users&apikey=%s", auth.TautulliURL, auth.TautulliKey), nil)
	req.Header.Add("Accept", "application/json")
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	log.WithFields(logrus.Fields{
		"request": "upstream",
		"api":     "tautulli",
		"took":    time.Since(start).Nanoseconds(),
	}).Debug("tautulli api call successful")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	tautulliResponse := TautulliResponse{}
	if err := json.Unmarshal(body, &tautulliResponse); err != nil {
		log.Error(err)
	}
	env.Config.Database.InsertTautulliUsers(tautulliResponse.Response.Data)
}

func syncOrganizr(env *Env) {
	log := env.Log
	client := &http.Client{}
	auth := env.Config.Authentication
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api?v1/user/list", auth.OrganizrURL), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("token", auth.OrganizrToken)
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	log.WithFields(logrus.Fields{
		"request": "upstream",
		"api":     "organizr",
		"took":    time.Since(start).Nanoseconds(),
	}).Debug("organizr api call successful")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	organizrResponse := OrganizrResponse{}
	if err := json.Unmarshal(body, &organizrResponse); err != nil {
		log.Error(err)
	}
	env.Config.Database.InsertOrganizrUsers(organizrResponse.Data.Users)
}
