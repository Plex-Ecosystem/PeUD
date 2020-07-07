package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

func ListUsers(env *Env, w http.ResponseWriter, r *http.Request) {
	/**
	* @api {get} /user Get Users
	* @apiName GetUser
	* @apiGroup Users
	* @apiDescription Get a list of users
	* @apiVersion 1.0.0
	* @apiExample {python3} python3
	*     import requests
	*     r = requests.get("http://127.0.0.1:8888/api/v1/user")
	*     r.json()
	* @apiExample {curl} curl
	*     curl -i http://127.0.0.1:8888/api/v1/user
	* @apiSuccess (200) {Object[]} users List of users
	* @apiSuccess (200) {Integer}  users.id Plex User ID
	* @apiSuccess (200) {String}   users.title Plex User Title
	* @apiSuccess (200) {String}   users.username Plex Username
	* @apiSuccess (200) {String}   users.email Plex email address
	* @apiSuccess (200) {String}   users.thumb Plex Thumbnail URL
	* @apiSuccess (200) {Boolean}  users.home Plex home member
	* @apiSuccess (200) {Boolean}  users.allowTuners Plex allow tuners (permission)
	* @apiSuccess (200) {Boolean}  users.allowSync Plex allow sync (permission)
	* @apiSuccess (200) {Boolean}  users.allowCameraUpload Plex allow camera upload (permission)
	* @apiSuccess (200) {Boolean}  users.allowChannels Plex allow using channels (permission)
	* @apiSuccess (200) {Boolean}  users.allowSubtitleAdmin Plex allow subtitle administration (permission)
	* @apiSuccessExample Success
	*     HTTP/1.1 200 OK
	*     [
	*       {
	*         "id": 12345678,
	*         "title": "plexuser1",
	*         "username": "plexuser1",
	*         "email": "plexuser1@gmail.com",
	*         "thumb": "https://plex.tv/users/a9s0df87a09sdf78/avatar?c=123412341234",
	*         "home": false,
	*         "allowTuners": false,
	*         "allowSync": false,
	*         "allowCameraUpload": false,
	*         "allowChannels": false,
	*         "allowSubtitleAdmin": false
	*       }
	*     ]
	* @apiError (500) {Object} response Response object
	* @apiError (500) {String} response.error Internal server error
	* @apiErrorExample {json} Internal Server Error
	* HTTP/1.1 500 Internal Server Error
	*     {
	*       "error": "InternalServerError"
	*     }
	* @apiSampleRequest /version
	 */
	env.toJSON(w, r, env.Config.Database.List())
}

func CreateUsers(env *Env, w http.ResponseWriter, r *http.Request) {
	logFields := logrus.Fields{
		"function": "CreateUsers",
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		env.Log.WithFields(logFields).Error(err)
	}
	isArray, isObject := arrayOrObject(resp)
	userList := make([]v1.PlexUser, 0)
	if isObject {
		var user v1.PlexUser
		if err := json.Unmarshal(resp, &user); err != nil {
			env.Log.WithFields(logFields).Error(err)
		}
		userList = append(userList, user)
	} else if isArray {
		if err := json.Unmarshal(resp, &userList); err != nil {
			env.Log.WithFields(logFields).Error(err)
		}
	}
	if err := env.Config.Database.Add(userList); err != nil {
		env.Log.WithFields(logFields).Error("Could not add users", err)
		http.Error(w, "Could not add users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Success")); err != nil {
		env.Log.WithFields(logFields).Error("Could not reply to your request", err)
		http.Error(w, "could not reply to your request", http.StatusInternalServerError)
	}
}
