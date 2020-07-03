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
	/**
	* @api {get} /version Get Version
	* @apiName Version
	* @apiGroup General
	* @apiDescription Get the build version and date information from the server
	* @apiVersion 1.0.0
	* @apiExample {python3} python3
	*     import requests
	*     r = requests.get("http://127.0.0.1:8888/api/v1/version")
	*     r.json()
	* @apiExample {curl} curl
	*     curl -i http://127.0.0.1:8888/api/v1/version
	* @apiSuccess (200) {Object} response Response object
	* @apiSuccess (200) {String} response.date Date compiled
	* @apiSuccess (200) {String} response.version Version number
	* @apiSuccessExample Success
	*     HTTP/1.1 200 OK
	*     {
	*       "date": "2020-07-03T07:01:59Z",
	*       "version": "0.1.0"
	*     }
	* @apiError (500) {Object} response Response object
	* @apiError (500) {String} response.error Internal server error
	* @apiErrorExample {json} Internal Server Error
	* HTTP/1.1 500 Internal Server Error
	*     {
	*       "error": "InternalServerError"
	*     }
	* @apiSampleRequest /version
	 */
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
