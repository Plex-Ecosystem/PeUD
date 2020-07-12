package handlers

import (
	"net/http"
)

// GetVersion returns a JSON formatted VersionHandler struct
func Version(env *Env, w http.ResponseWriter, r *http.Request) {
	env.toJSON(w, r, env.Build)
}
