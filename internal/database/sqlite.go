package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

type Database struct {
	Name string
	Log  *logrus.Logger
	Type string
	DB   *sql.DB
}

func (d *Database) Init() {
	logFields := logrus.Fields{
		"struct":   "Database",
		"function": "Init",
	}
	if d.Type == "sqlite" {
		db, err := sql.Open("sqlite3", d.Name)
		if err != nil {
			d.Log.WithFields(logFields).Error(err)
		}

		createPlexUserTable := `CREATE TABLE IF NOT EXISTS plex_users (
                                  id INTEGER PRIMARY KEY,
                                  title TEXT NOT NULL,
                                  username TEXT NOT NULL UNIQUE,
                                  email TEXT NOT NULL UNIQUE,
                                  thumb TEXT NOT NULL,
                                  home BOOL NOT NULL,
                                  allowTuners BOOL NOT NULL,
                                  allowSync BOOL NOT NULL,
                                  allowCameraUpload BOOL NOT NULL,
                                  allowChannels BOOL NOT NULL,
                                  allowSubtitleAdmin BOOL NOT NULL
                                );`
		stmt, err := db.Prepare(createPlexUserTable)
		if err != nil {
			d.Log.WithFields(logFields).Error(err)
		}

		_, err = stmt.Exec()
		if err != nil {
			d.Log.WithFields(logFields).Error(err)
		}

		d.DB = db
	}

}

func (d *Database) List() []v1.PlexUser {
	logFields := logrus.Fields{
		"struct":   "Database",
		"function": "List",
	}
	db := d.DB
	rows, err := db.Query("SELECT * FROM plex_users")
	if err != nil {
		d.Log.WithFields(logFields).Error(err)
	}
	plexUserList := make([]v1.PlexUser, 0)
	for rows.Next() {
		user := v1.PlexUser{
			Title:              "",
			Username:           "",
			Email:              "",
			Thumb:              "",
			Home:               false,
			AllowTuners:        false,
			AllowSync:          false,
			AllowCameraUpload:  false,
			AllowChannels:      false,
			AllowSubtitleAdmin: false,
		}
		err = rows.Scan(
			&user.ID, &user.Title, &user.Username, &user.Email, &user.Thumb, &user.Home, &user.AllowTuners,
			&user.AllowSync, &user.AllowCameraUpload, &user.AllowChannels, &user.AllowSubtitleAdmin,
		)
		plexUserList = append(plexUserList, user)
	}
	if err := rows.Close(); err != nil {
		d.Log.WithFields(logFields).Error(err)
	}
	return plexUserList
}

func (d *Database) Insert(userList []v1.PlexUser) error {
	logFields := logrus.Fields{
		"struct":   "Database",
		"function": "Insert",
	}
	db := d.DB
	insertStatement := `INSERT INTO plex_users(
	                      id, title, username, email, thumb, home, allowTuners, allowSync,
                          allowCameraUpload, allowChannels, allowSubtitleAdmin
                        )  values(?,?,?,?,?,?,?,?,?,?,?)`
	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		d.Log.WithFields(logFields).Error(err)
	}
	for _, user := range userList {
		_, err := stmt.Exec(
			&user.ID, &user.Title, &user.Username, &user.Email, &user.Thumb, &user.Home,
			&user.AllowTuners, &user.AllowSync, &user.AllowCameraUpload, &user.AllowChannels, &user.AllowSubtitleAdmin,
		)
		if err != nil {
			d.Log.WithFields(logFields).Error(err)
		}
	}
	return nil
}
