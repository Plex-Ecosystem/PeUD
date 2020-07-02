package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

func ListUsers(env *Env, w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users Users
	// Returns a list of existing users.
	//     Responses:
	//       200:
	//         Description: Successfully returned the list of users
	//       500:
	//         Description: InternalServerError
	logFields := logrus.Fields{
		"function": "ListUsers",
	}
	plexUserList := env.Config.Database.List()
	responseJSON, err := json.Marshal(plexUserList)
	if err != nil {
		env.Log.WithFields(logFields).Error("Could not marshal json body", err)
		http.Error(w, "Could not marshal your request", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responseJSON); err != nil {
		env.Log.WithFields(logFields).Error("Could not reply to your request", err)
		http.Error(w, "could not reply to your request", http.StatusInternalServerError)
	}
}

func CreateUsers(env *Env, w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /users Users
	// Creates users.
	//     Responses:
	//       201:
	//         Description: Successfully created the user(s)
	//       400:
	//         Description: Invalid user input
	//       500:
	//         Description: InternalServerError
	logFields := logrus.Fields{
		"function": "CreateUsers",
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print("asdf")
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
	if err := env.Config.Database.Insert(userList); err != nil {
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
