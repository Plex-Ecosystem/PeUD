package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
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

func (d *Database) ListUsers(endpoint string) interface{} {
	log := d.Log.WithField("function", "ListUsers")
	rows, err := d.Db.Query(fmt.Sprintf("SELECT * FROM %sUsers", endpoint))
	if err != nil {
		log.Error(err)
	}
	switch endpoint {
	case "plex":
		users := make([]v1.PlexUser, 0)
		sqlx.StructScan(rows, &users)
		return users
	case "tautulli":
		users := make([]v1.TautulliUser, 0)
		sqlx.StructScan(rows, &users)
		return users
	case "organizr":
		users := make([]v1.OrganizrUser, 0)
		sqlx.StructScan(rows, &users)
		return users
	case "ombi":
		users := make([]v1.OmbiUser, 0)
		sqlx.StructScan(rows, &users)
		return users
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

func (d *Database) dropRows(table string) {
	log := d.Log.WithField("function", "dropRows")
	if _, err := d.Exec(fmt.Sprintf("DELETE FROM %s", table)); err != nil {
		log.Error(err)
		return
	}
	log.Debug("dropped all rows in ", table)
}
