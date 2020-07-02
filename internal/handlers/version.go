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
	// swagger:route GET /version version
	// Outputs the build version and date.
	// Requires authentication
	//     Responses:
	//       200:
	//         Description: Success
	//       500:
	//         Description: InternalServerError
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
