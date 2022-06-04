package data

import (
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"week04/app/service/store/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, ConnectDB, NewStoreRepo)

type Data struct {
	db  *sql.DB
	log *log.Helper
}

func NewData(data *conf.Data, db *sql.DB, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{db: db, log: l}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func ConnectDB(config *conf.Data, logger log.Logger) *sql.DB {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data/ent"))

	dbconn, er := sql.Open(config.Database.Driver, config.Database.Source)
	if er != nil {
		log.Fatalf("Connect Mysql Error Err: %s", er.Error())
	}
	if config.Database.MaxOpenConns != nil {
		dbconn.SetMaxOpenConns(*config.Database.MaxOpenConns)
	}

	if config.Database.ConnMaxLifeTime != nil {
		dbconn.SetConnMaxLifetime(*config.Database.ConnMaxLifeTime)
	}

	if err := dbconn.Ping(); err != nil {
		dbconn.Close()
		log.Fatalf("Connect Mysql Error Err: %s", err.Error())
	}
	return dbconn
}
