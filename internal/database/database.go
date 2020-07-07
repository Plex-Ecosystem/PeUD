package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/modl"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
)

type Database struct {
	Name string
	Type string
	Log  *logrus.Entry
	*modl.DbMap
}

func fixColMap(t *modl.TableMap, s interface{}) {
	v := reflect.TypeOf(s)
	switch v.Name() {
	case "PlexUser":
		t.SetKeys(false, "id")
	case "TautulliUser":
		t.SetKeys(false, "rowid")
	case "OrganizrUser":
		t.SetKeys(false, "id")
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := field.Tag.Get("peud")
		cm := t.ColMap(strings.ToLower(field.Name))
		if strings.Contains(tag, "u") {
			cm.Unique = true
		}
	}
}

func (d *Database) setDialect() {
	switch d.Type {
	case "mysql":
		// TODO: Implement MySQL
		// d.Dialect = &modl.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}
	case "postgres":
		// TODO: Implement Postgres
		// d.Dialect = &modl.PostgresDialect{}
	case "sqlite3":
		d.Dialect = &modl.SqliteDialect{}
	}
}

func (d *Database) buildTables(tables []interface{}) {
	log := d.Log.WithField("function", "buildTables")
	for _, i := range tables {
		fixColMap(d.AddTableWithName(i, strcase.ToLowerCamel(fmt.Sprintf("%ss", reflect.TypeOf(i).Name()))), i)
		if err := d.CreateTablesIfNotExists(); err != nil {
			log.Error(err)
		}
	}
}

func (d *Database) Init() {
	log := d.Log.WithField("function", "init")
	d.setDialect()
	var err error
	d.Db, err = sql.Open(d.Type, d.Name)
	d.DbMap = modl.NewDbMap(d.Db, d.Dialect)
	if err != nil {
		log.Fatal(err)
	}
	// create tables
	tables := []interface{}{
		v1.PlexUser{},
		v1.TautulliUser{},
		v1.OrganizrUser{},
		v1.OmbiUser{},
	}
	d.buildTables(tables)
}

func (d *Database) ListPlexUsers() []*v1.PlexUser {
	log := d.Log.WithField("function", "list")
	users := make([]*v1.PlexUser, 0)
	if err := d.Select(&users, "SELECT * FROM plexUsers"); err != nil {
		log.Error(err)
	}
	return users
}

func (d *Database) ListTautulliUsers() []*v1.TautulliUser {
	log := d.Log.WithField("function", "list")
	users := make([]*v1.TautulliUser, 0)
	if err := d.Select(&users, "SELECT * FROM tautulliUsers"); err != nil {
		log.Error(err)
	}
	return users
}

func (d *Database) ListOrganizrUsers() []*v1.OrganizrUser {
	log := d.Log.WithField("function", "list")
	users := make([]*v1.OrganizrUser, 0)
	if err := d.Select(&users, "SELECT * FROM organizrUsers"); err != nil {
		log.Error(err)
	}
	return users
}

func (d *Database) ListOmbiUsers() []*v1.OmbiUser {
	log := d.Log.WithField("function", "list")
	users := make([]*v1.OmbiUser, 0)
	if err := d.Select(&users, "SELECT * FROM ombiUsers"); err != nil {
		log.Error(err)
	}
	return users
}

func (d *Database) dropRows(table string) {
	log := d.Log.WithField("function", "dropRows")
	if _, err := d.Exec(fmt.Sprintf("DELETE FROM %s", table)); err != nil {
		log.Error(err)
		return
	}
	log.Debug("dropped all rows in ", table)
}

func (d *Database) InsertPlexUsers(userList []v1.PlexUser) error {
	log := d.Log.WithField("function", "add")
	d.dropRows("plexUsers")
	for _, user := range userList {
		if user.Username == "Local" {
			continue
		}
		if err := d.Insert(&user); err != nil {
			log.Error(err)
		}
	}
	log.Debugf("added users to plexUsers")
	return nil
}

func (d *Database) InsertTautulliUsers(userList []v1.TautulliUser) error {
	log := d.Log.WithField("function", "add")
	d.dropRows("tautulliUsers")
	for _, user := range userList {
		if err := d.Insert(&user); err != nil {
			log.Error(err)
		}
	}
	log.Debugf("added users to tautulliUsers")
	return nil
}

func (d *Database) InsertOrganizrUsers(userList []v1.OrganizrUser) error {
	log := d.Log.WithField("function", "add")
	d.dropRows("organizrUsers")
	for _, user := range userList {
		if err := d.Insert(&user); err != nil {
			log.Error(err)
		}
	}
	log.Debugf("added users to organizrUsers")
	return nil
}

func (d *Database) InsertOmbiUsers(userList []v1.OmbiUser) error {
	log := d.Log.WithField("function", "add")
	d.dropRows("ombiUsers")
	for _, user := range userList {
		if err := d.Insert(&user); err != nil {
			log.Error(err)
		}
	}
	log.Debugf("added users to ombiUsers")
	return nil
}
