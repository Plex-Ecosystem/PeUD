package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx/reflectx"
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
	if v.Name() == "PlexUser" {
		t.SetKeys(false, "id")
	} else if v.Name() == "TautulliUser" {
		t.SetKeys(false, "userID")
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := field.Tag.Get("peud")
		cm := t.ColMap(strings.ToLower(field.Name))
		cm.ColumnName = strcase.ToLowerCamel(field.Name)
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
	modl.TableNameMapper = strcase.ToLowerCamel
	d.DbMap = modl.NewDbMap(d.Db, d.Dialect)
	if err != nil {
		log.Fatal(err)
	}
	// reuse json tags to map to structs
	d.Dbx.Mapper = reflectx.NewMapperFunc("json", strcase.ToLowerCamel)
	// create tables
	tables := []interface{}{
		v1.PlexUser{},
		v1.TautulliUser{},
	}
	d.buildTables(tables)
}

func (d *Database) List() []*v1.PlexUser {
	log := d.Log.WithField("function", "list")
	var plexUsers []*v1.PlexUser
	if err := d.Select(&plexUsers, "SELECT * FROM plexUser"); err != nil {
		log.Error(err)
	}
	return plexUsers
}

func (d *Database) Add(userList []v1.PlexUser) error {
	log := d.Log.WithField("function", "add")
	for _, user := range userList {
		if err := d.Insert(&user); err != nil {
			log.Error(err)
		}
		log.Debugf("Successfully created new user: %s", user.Username)
	}
	return nil
}
