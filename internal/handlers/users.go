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
	// swagger:operation GET  /users listUsers
	// Returns a list of existing users
	// ---
	// tags:
	//   - Users
	// produces:
	//   - application/json
	// responses:
	//   '200':
	//     description: Success
	//     schema:
	//       type: array
	//       items:
	//         properties:
	//           id:
	//             type: integer
	//             description: Plex ID
	//           title:
	//             type: string
	//             description: Plex User Title
	//           username:
	//             type: string
	//             description: Plex Username
	//           email:
	//             type: string
	//             description: Plex email
	//           thumb:
	//             type: string
	//             description: Plex Thumbnail URL
	//           home:
	//             type: bool
	//             description: Plex Home User
	//           allowTuners:
	//             type: boolean
	//             description: Allow Tuners
	//           allowSync:
	//             type: boolean
	//             description: Allow Sync
	//           allowCameraUpload:
	//             type: boolean
	//             description: Allow Camera Upload
	//           allowChannels:
	//             type: boolean
	//             description: Allow Using Channels
	//           allowSubtitleAdmin:
	//             type: boolean
	//             description: Allow Subtitle Admin
	//       example:
	//         - id: 123123123
	//           title: user1
	//           username: user1
	//           email: user1@gmail.com
	//           thumb: https://plex.tv/users/asdfasdf/avatar?c=asdfasdf
	//           home: false
	//           allowTuners: false
	//           allowSync: false
	//           allowCameraUpload: false
	//           allowChannels: false
	//           allowSubtitleAdmin: false
	//         - id: 234234234
	//           title: user2
	//           username: user2
	//           email: user1@gmail.com
	//           thumb: https://plex.tv/users/asdfasdf/avatar?c=asdfasdf
	//           home: false
	//           allowTuners: false
	//           allowSync: false
	//           allowCameraUpload: false
	//           allowChannels: false
	//           allowSubtitleAdmin: false
	//   '500':
	//      description: InternalServerError
	// x-codeSamples:
	//   - lang: python3
	//     source: requests.get("http://127.0.0.1:8888/api/v1/users")
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
	// swagger:operation POST  /users createUsers
	// Create new users
	// ---
	// tags:
	//   - Users
	// produces:
	//   - application/json
	// responses:
	//   '201':
	//     description: Created
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
