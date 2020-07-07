package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

func Sync(env *Env, w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("only")
	if strings.Contains(filter, "plex") {
		syncPlex(env)
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
