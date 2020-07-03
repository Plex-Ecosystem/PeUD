package handlers

import (
	"encoding/json"
	"net/http"
)

type VersionHandler struct {
	Date    *string `json:"date"`
	Version *string `json:"version"`
}

func NewVersionHandler(version *string, date *string) *VersionHandler {
	return &VersionHandler{
		Date:    date,
		Version: version,
	}
}

// GetVersion returns a JSON formatted VersionHandler struct
func (v *VersionHandler) GetVersion(handlerEnv *Env, w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /version Version
	// Get server version and build date
	// ---
	// tags:
	//   - General
	// produces:
	//   - application/json
	// responses:
	//   '200':
	//     description: Success
	//     schema:
	//       type: object
	//       properties:
	//         date:
	//           type: string
	//           description: Server build date
	//         version:
	//           type: string
	//           description: Server version
	//       example:
	//         date: 2020-07-03T07:01:59Z
	//         version: 1.0.0
	//
	// x-codeSamples:
	//   - lang: python3
	//     source: requests.get("http://127.0.0.1:8888/api/v1/version")
	log := handlerEnv.Log
	b, err := json.Marshal(v)
	if err != nil {
		log.WithField("handler", "version").Error("could not marshal app version", err)
		http.Error(w, "Could not get app version", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.WithField("handler", "version").Error("could not write response", err)
		http.Error(w, "Could not write http response", http.StatusInternalServerError)
	}
}
