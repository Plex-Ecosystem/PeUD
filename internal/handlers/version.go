package handlers

import (
	"net/http"
)

// GetVersion returns a JSON formatted VersionHandler struct
func Version(Env *Env, w http.ResponseWriter, r *http.Request) {
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
	//log := Env.Log
	Env.toJSON(w, r, Env.Build)
}
