package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func ListUsers(env *Env, w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	endpoint := splitPath[len(splitPath)-1]
	env.toJSON(w, r, env.Config.Database.ListUsers(endpoint, r.URL.Query()))
}

func GetUser(env *Env, w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	endpoint := splitPath[len(splitPath)-2]
	userID := chi.URLParam(r, "id")
	env.toJSON(w, r, env.Config.Database.GetUser(endpoint, userID))

}

//func CreateUsers(env *Env, w http.ResponseWriter, r *http.Request) {
//	logFields := logrus.Fields{
//		"function": "CreateUsers",
//	}
//	resp, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		env.Log.WithFields(logFields).Error(err)
//	}
//	isArray, isObject := arrayOrObject(resp)
//	userList := make([]v1.PlexUser, 0)
//	if isObject {
//		var user v1.PlexUser
//		if err := json.Unmarshal(resp, &user); err != nil {
//			env.Log.WithFields(logFields).Error(err)
//		}
//		userList = append(userList, user)
//	} else if isArray {
//		if err := json.Unmarshal(resp, &userList); err != nil {
//			env.Log.WithFields(logFields).Error(err)
//		}
//	}
//	if err := env.Config.Database.InsertPlexUsers(userList); err != nil {
//		env.Log.WithFields(logFields).Error("Could not add users", err)
//		http.Error(w, "Could not add users", http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusCreated)
//	if _, err := w.Write([]byte("Success")); err != nil {
//		env.Log.WithFields(logFields).Error("Could not reply to your request", err)
//		http.Error(w, "could not reply to your request", http.StatusInternalServerError)
//	}
//}

// Beginning of API Documentation
/**
@api {get} /users/plex Plex
@apiGroup Users
@apiDescription Get a list of Plex users
@apiVersion 1.0.0
@apiParam {String}  [id]        Plex user ID
@apiParam {String}  [uuid]      Plex UUID
@apiParam {String}  [username]  Plex Username
@apiParam {String}  [email]     Plex email address
@apiParam {Boolean} [home]      Plex home member
@apiParam {String}  [status]    Plex friend status
@apiParam {Boolean} [admin]     Plex is user an admin
@apiParam {Boolean} [protected] Plex is user protected
@apiParamExample {python3} python3
    import requests
    params = {
        'status': 'accepted',
        'protected': True
    }
    response = requests.get("http://127.0.0.1:8888/api/v1/users/plex", params=params)
    response.json()
@apiParamExample {bash} curl
    curl -i http://127.0.0.1:8888/api/v1/users/plex
@apiExample {python3} python3
    import requests
    response = requests.get("http://127.0.0.1:8888/api/v1/users/plex")
    response.json()
@apiExample {curl} curl
    curl -i http://127.0.0.1:8888/api/v1/users/plex
@apiSuccess {List}    plexUsers                      List of plexUser objects
@apiSuccess {Object}  plexUsers.plexUser             plexUser object
@apiSuccess {Integer} plexUsers.plexUser.id          Plex user ID
@apiSuccess {String}  plexUsers.plexUser.uuid        Plex UUID
@apiSuccess {Boolean} plexUsers.plexUser.hasPassword Plex Password check
@apiSuccess {String}  plexUsers.plexUser.username    Plex Username
@apiSuccess {String}  plexUsers.plexUser.email       Plex email address
@apiSuccess {String}  plexUsers.plexUser.thumb       Plex Thumbnail URL
@apiSuccess {String}  plexUsers.plexUser.title       Plex User Title
@apiSuccess {Boolean} plexUsers.plexUser.home        Plex home member
@apiSuccess {Boolean} plexUsers.plexUser.restricted  Plex is user restricted
@apiSuccess {String}  plexUsers.plexUser.status      Plex friend status
@apiSuccess {Boolean} plexUsers.plexUser.admin       Plex is user an admin
@apiSuccess {Boolean} plexUsers.plexUser.guest       Plex is user a guest
@apiSuccess {Boolean} plexUsers.plexUser.protected   Plex is user protected
@apiSuccessExample Success
    HTTP/1.1 200 OK
    [
      {
        "id": 12345678,
        "uuid": "0986fg6dh0786u0fgh",
        "hasPassword": false,
        "username": "User1",
        "email": "plexuser1@gmail.com",
        "thumb": "https://plex.tv/users/a9s0df87a09sdf78/avatar?c=123412341234",
        "title": "User1",
        "home": false,
        "restricted": false,
        "status": "accepted",
        "admin": false,
        "guest": false,
        "protected": false
      }
    ]
@apiUse InternalServerError
*/
/**
@api {get} /users/tautulli Tautulli
@apiGroup Users
@apiDescription Get a list of Tautulli users
@apiVersion 1.0.0
@apiExample {python3} python3
    import requests
    r = requests.get("http://127.0.0.1:8888/api/v1/users/tautulli")
    r.json()
@apiExample {curl} curl
    curl -i http://127.0.0.1:8888/api/v1/users/tautulli
@apiSuccess (200) {List}    tautulliUsers                                List of tautulliUser objects
@apiSuccess (200) {Object}  tautulliUsers.tautulliUser                   tautulliUser object
@apiSuccess (200) {Integer} tautulliUsers.tautulliUser.row_id            Tautulli database row ID
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.username          Plex username
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.email             Plex email address
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.thumb             Plex Thumbnail URL
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.filter_all        Plex filter for everything
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.filter_movies     Plex filter for movies
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.filter_music      Plex filter for music
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.filter_photos     Plex filter for photos
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.filter_television Plex filter for television
@apiSuccess (200) {Integer} tautulliUsers.tautulliUser.user_id           Plex user ID
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.friendly_name     Tautulli friendly name
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.is_active         Tautulli check for Plex friend status
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.is_admin          Plex is user an admin
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.is_home_user      Plex home member
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.is_allow_sync     Plex allowed to sync
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.is_restricted     Plex is user restricted
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.do_notify         Tautulli receive notifications
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.keep_history      Tautulli keep watch history
@apiSuccess (200) {Boolean} tautulliUsers.tautulliUser.allow_guest       Tautulli allow guest access
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.server_token      Plex server token
@apiSuccess (200) {String}  tautulliUsers.tautulliUser.shared_libraries  Plex shared libraries

@apiSuccessExample Success
    HTTP/1.1 200 OK
    [
      {
        "row_id": 2,
        "username": "User1",
        "email": "plexuser1@gmail.com",
        "thumb": "https://plex.tv/users/a9s0df87a09sdf78/avatar?c=123412341234",
        "filter_all": "",
        "filter_movies": "",
        "filter_music": "",
        "filter_photos": "",
        "filter_television": "",
        "user_id": 12345678,
        "friendly_name": "John Doe",
        "is_active": true,
        "is_admin": true,
        "is_home_user": true,
        "is_allow_sync": true,
        "is_restricted": false,
        "do_notify": true,
        "keep_history": true,
        "allow_guest": false,
        "server_token": "9a8sd7fa3sd9780dfa",
        "shared_libraries": "8;5;4;7;2"
      }
    ]
@apiUse InternalServerError
*/
/**
@api {get} /users/ombi Ombi
@apiGroup Users
@apiDescription Get a list of Ombi users
@apiVersion 1.0.0
@apiExample {python3} python3
    import requests
    r = requests.get("http://127.0.0.1:8888/api/v1/users/ombi")
    r.json()
@apiExample {curl} curl
    curl -i http://127.0.0.1:8888/api/v1/users/ombi
@apiSuccess (200) {List}    ombiUsers                            List of tautulliUser objects
@apiSuccess (200) {Object}  ombiUsers.ombiUser                   tautulliUser object
@apiSuccess (200) {Integer} ombiUsers.ombiUser.row_id            Tautulli database row ID
@apiSuccess (200) {String}  ombiUsers.ombiUser.username          Plex username
@apiSuccess (200) {String}  ombiUsers.ombiUser.email             Plex email address
@apiSuccess (200) {String}  ombiUsers.ombiUser.thumb             Plex Thumbnail URL
@apiSuccess (200) {String}  ombiUsers.ombiUser.filter_all        Plex filter for everything
@apiSuccess (200) {String}  ombiUsers.ombiUser.filter_movies     Plex filter for movies
@apiSuccess (200) {String}  ombiUsers.ombiUser.filter_music      Plex filter for music
@apiSuccess (200) {String}  ombiUsers.ombiUser.filter_photos     Plex filter for photos
@apiSuccess (200) {String}  ombiUsers.ombiUser.filter_television Plex filter for television
@apiSuccess (200) {Integer} ombiUsers.ombiUser.user_id           Plex user ID
@apiSuccess (200) {String}  ombiUsers.ombiUser.friendly_name     Tautulli friendly name
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.is_active         Tautulli check for Plex friend status
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.is_admin          Plex is user an admin
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.is_home_user      Plex home member
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.is_allow_sync     Plex allowed to sync
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.is_restricted     Plex is user restricted
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.do_notify         Tautulli receive notifications
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.keep_history      Tautulli keep watch history
@apiSuccess (200) {Boolean} ombiUsers.ombiUser.allow_guest       Tautulli allow guest access
@apiSuccess (200) {String}  ombiUsers.ombiUser.server_token      Plex server token
@apiSuccess (200) {String}  ombiUsers.ombiUser.shared_libraries  Plex shared libraries

@apiSuccessExample Success
    HTTP/1.1 200 OK
    [
      {
        "row_id": 2,
        "username": "User1",
        "email": "plexuser1@gmail.com",
        "thumb": "https://plex.tv/users/a9s0df87a09sdf78/avatar?c=123412341234",
        "filter_all": "",
        "filter_movies": "",
        "filter_music": "",
        "filter_photos": "",
        "filter_television": "",
        "user_id": 12345678,
        "friendly_name": "John Doe",
        "is_active": true,
        "is_admin": true,
        "is_home_user": true,
        "is_allow_sync": true,
        "is_restricted": false,
        "do_notify": true,
        "keep_history": true,
        "allow_guest": false,
        "server_token": "9a8sd7fa3sd9780dfa",
        "shared_libraries": "8;5;4;7;2"
      }
    ]
@apiUse InternalServerError
*/
