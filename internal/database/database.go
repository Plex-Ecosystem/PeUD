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
	case "TautulliUser":
		t.SetKeys(false, "rowid")
	case "OmbiUserQualityProfile", "OmbiUserClaim":
		t.SetKeys(false, "userid")
	default:
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
		v1.PlexUserServer{},
		v1.TautulliUser{},
		v1.OrganizrUser{},
		v1.OmbiUser{},
		v1.OmbiUserClaim{},
		v1.OmbiUserQualityProfile{},
	}
	d.buildTables(tables)
}

func (d *Database) dropRows(table string) {
	log := d.Log.WithField("function", "dropRows")
	if _, err := d.Exec(fmt.Sprintf("DELETE FROM %s", table)); err != nil {
		log.Error(err)
		return
	}
	log.Debug("dropped all rows in ", table)
}
