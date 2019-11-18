package initDB

import (
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
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
