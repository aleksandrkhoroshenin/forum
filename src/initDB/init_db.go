package initDB

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"log"
)

type DbManager interface {
	DbConnect(host, database, user, password string, port uint16) error
	GetConnPool() *pgx.ConnPool
}

type InitDB struct {
	pool *pgx.ConnPool
}

func Init() DbManager {
	return &InitDB{}
}

func (db InitDB) GetConnPool() *pgx.ConnPool {
	return db.pool
}

func (db *InitDB) DbConnect(host, database, user, password string, port uint16) (err error) {
	l := logrus.New()
	l.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	logger := logrusadapter.NewLogger(l)
	runtimeParams := make(map[string]string)
	runtimeParams["application_name"] = "dz"
	conConfig := pgx.ConnConfig{
		Host:           host,
		Port:           port,
		Database:       database,
		User:           user,
		Password:       password,
		TLSConfig:      nil,
		UseFallbackTLS: false,
		RuntimeParams:  runtimeParams,
		LogLevel:       5,
		Logger:         logger,
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     conConfig,
		MaxConnections: 20,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}
	p, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		log.Println(err)
		return err
	}
	db.pool = p
	return nil
}
