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
