package initDB

import (
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"log"
)

type DbManager interface {
	DbConnect(host, port, database, user, password string) error
}
type InitDB struct {
	pool *pgx.ConnPool
}

func Init() DbManager {
	return &InitDB{}
}

func (db *InitDB) DbConnect(host, port, database, user, password string) (err error) {
	runtimeParams := make(map[string]string)
	runtimeParams["application_name"] = "dz"
	conConfig := pgx.ConnConfig{
		Host:           host,
		Port:           5432,
		Database:       "docker",
		User:           "docker",
		Password:       "docker",
		TLSConfig:      nil,
		UseFallbackTLS: false,
		RuntimeParams:  runtimeParams,
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
