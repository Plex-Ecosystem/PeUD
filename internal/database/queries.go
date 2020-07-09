package database

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

func cleanOpts(p []string) []string {
	params := make([]string, 0)
	if len(p) != 0 {
		for _, i := range p {
			param := strings.TrimSpace(i)
			if param != "" {
				params = append(params, fmt.Sprintf("'%s'", param))
			}
		}
	}
	return params
}

func isInSlice(slice []string, str string) bool {
	for _, i := range slice {
		if i == str {
			return true
		}
	}
	return false
}

func makeSQLConditions(validConditions []string, urlQuery url.Values) []string {
	conditions := make([]string, 0)
	for param, opts := range urlQuery {
		var vals []string
		if isInSlice(validConditions, param) {
			if vals = cleanOpts(opts); len(vals) != 0 {
				conditions = append(conditions, fmt.Sprintf("%s in (%s)", param, strings.Join(vals, ",")))
			}
		}
	}
	return conditions
}

func (d *Database) ListUsers(endpoint string, urlQuery url.Values) interface{} {
	log := d.Log.WithField("function", "ListUsers")
	queryOpts := []string{fmt.Sprintf("SELECT * FROM %sUsers", endpoint)}
	sqlQuery := queryOpts[0]
	log.Tracef("base sql query is '%s'", sqlQuery)
	if len(urlQuery) != 0 {
		log.Trace("url query params detected. checking conditions")
		conditions := makeSQLConditions(v1.PlexParams(), urlQuery)
		if len(conditions) != 0 {
			log.Tracef("valid sql conditions are %s", conditions)
			queryOpts = append(queryOpts, "WHERE", strings.Join(conditions, " AND "))
			sqlQuery = strings.Join(queryOpts, " ")
		} else {
			log.Tracef("No valid query params. requested: %s", urlQuery)
		}
	}
	log.Tracef("full sql query is '%s'", sqlQuery)
	rows, err := d.Db.Query(sqlQuery)
	if err != nil {
		log.Error(err)
	}
	switch endpoint {
	case "plex":
		users := make([]v1.PlexUser, 0)
		if err := sqlx.StructScan(rows, &users); err != nil {
			log.Error(err)
		}
		return users
	case "tautulli":
		users := make([]v1.TautulliUser, 0)
		if err := sqlx.StructScan(rows, &users); err != nil {
			log.Error(err)
		}
		return users
	case "organizr":
		users := make([]v1.OrganizrUser, 0)
		if err := sqlx.StructScan(rows, &users); err != nil {
			log.Error(err)
		}
		return users
	case "ombi":
		users := make([]v1.OmbiUser, 0)
		if err := sqlx.StructScan(rows, &users); err != nil {
			log.Error(err)
		}
		return users
	default:
		var placeholder interface{}
		return placeholder
	}
}

func (d *Database) GetUser(endpoint string, id string) interface{} {
	log := d.Log.WithField("function", "GetUser")
	switch endpoint {
	case "plex":
		user := make([]v1.PlexUser, 0)
		if err := d.Select(&user, "SELECT * from plexUsers WHERE id = ? LIMIT 1", id); err != nil {
			log.Error(err)
		}
		return user
	case "tautulli":
		var user []v1.TautulliUser
		if err := d.Select(&user, "SELECT * from tautulliUsers WHERE user_id = ? LIMIT 1", id); err != nil {
			log.Error(err)
		}
		return user
	case "organizr":
		var user []v1.OrganizrUser
		if err := d.Select(&user, "SELECT * from organizrUsers WHERE id = ? LIMIT 1", id); err != nil {
			log.Error(err)
		}
		return user
	case "ombi":
		var user []v1.OmbiUser
		if err := d.Select(&user, "SELECT * from ombiUsers WHERE id = ? LIMIT 1", id); err != nil {
			log.Error(err)
		}
		return user
	default:
		var placeholder interface{}
		return placeholder
	}
}

func (d *Database) InsertUsers(table string, v interface{}) error {
	log := d.Log.WithField("function", "InsertUsers")
	d.dropRows(table)
	switch x := v.(type) {
	case []v1.PlexUser:
		for _, user := range x {
			if err := d.Insert(&user); err != nil {
				log.Error(err)
			}
		}
	case []v1.OrganizrUser:
		for _, user := range x {
			if err := d.Insert(&user); err != nil {
				log.Error(err)
			}
		}
	case []v1.OmbiUser:
		for _, user := range x {
			if err := d.Insert(&user); err != nil {
				log.Error(err)
			}
		}
	case []v1.TautulliUser:
		for _, user := range x {
			if err := d.Insert(&user); err != nil {
				log.Error(err)
			}
		}
	}
	log.Debug("added users to ", table)
	return nil
}
