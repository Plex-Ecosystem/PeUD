package handlers

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

func Sync(env *Env, w http.ResponseWriter, r *http.Request) {
	filter := strings.Split(r.URL.Query().Get("only"), ",")
	if len(filter) == 0 {
		filter = []string{"organizr", "tautulli", "plex", "ombi"}
	}
	for _, i := range filter {
		sync(i, env)
	}
	w.WriteHeader(http.StatusNoContent)
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
	// TODO: switch this to a env option
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: t}
	var (
		auth = env.Config.Authentication
		db   = env.Config.Database
		h    map[string][]string
	)
	switch e {
	case "plex":
		u := "https://plex.tv/api/users"
		h = map[string][]string{
			"X-Plex-Client-Identifier": {"PeUD"},
			"X-Plex-Token":             {auth.PlexToken},
		}
		b := sharedRequest(c, u, h, l, e)
		plexResponse := v1.PlexResponse{}
		if err := xml.Unmarshal(b, &plexResponse); err != nil {
			l.Error(err)
		}
		if err := db.InsertUsers("plexUsers", plexResponse.PlexUsers); err != nil {
			l.Error(err)
		}
	case "tautulli":
		u := fmt.Sprintf("%s/api/v2?cmd=get_users&apikey=%s", auth.TautulliURL, auth.TautulliKey)
		b := sharedRequest(c, u, h, l, e)
		tautulliResponse := v1.TautulliResponse{}
		if err := json.Unmarshal(b, &tautulliResponse); err != nil {
			l.Error(err)
		}
		if err := db.InsertUsers("tautulliUsers", tautulliResponse.Response.Data); err != nil {
			l.Error(err)
		}
	case "organizr":
		u := fmt.Sprintf("%s/api?v1/user/list", auth.OrganizrURL)
		h = map[string][]string{"token": {auth.OrganizrToken}}
		b := sharedRequest(c, u, h, l, e)
		organizrResponse := v1.OrganizrResponse{}
		if err := json.Unmarshal(b, &organizrResponse); err != nil {
			l.Error(err)
		}
		if err := db.InsertUsers("organizrUsers", organizrResponse.Data.Users); err != nil {
			l.Error(err)
		}
	case "ombi":
		u := fmt.Sprintf("%s/api/v1/Identity/Users", auth.OmbiURL)
		h = map[string][]string{"ApiKey": {auth.OmbiKey}}
		b := sharedRequest(c, u, h, l, e)
		ombiUsers := make([]v1.OmbiUser, 0)
		if err := json.Unmarshal(b, &ombiUsers); err != nil {
			l.Error(err)
		}
		if err := db.InsertUsers("ombiUsers", ombiUsers); err != nil {
			l.Error(err)
		}
	}
}
