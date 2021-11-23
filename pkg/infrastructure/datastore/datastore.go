package datastore

import (
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"

	"github.com/go-sql-driver/mysql"
)

// NewDB returns database client with ORM
func NewDB() (*ent.Client, error) {
	DBMS := "mysql"
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	mc := mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
			"charset":   config.C.Database.Params.Charset,
			"loc":       config.C.Database.Params.Loc,
		},
	}

	return ent.Open(DBMS, mc.FormatDSN(), entOptions...)
}
