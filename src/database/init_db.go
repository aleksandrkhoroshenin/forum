package database

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"log"
)

const sqlInit = `
CREATE EXTENSION IF NOT EXISTS citext;

DROP TABLE IF EXISTS "forums" CASCADE;
DROP TABLE IF EXISTS "posts" CASCADE;
DROP TABLE IF EXISTS "threads" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "votes" CASCADE;
DROP TABLE IF EXISTS "forum_users" CASCADE;

CREATE TABLE IF NOT EXISTS users (
                                     "nickname" CITEXT UNIQUE PRIMARY KEY,
                                     "email"    CITEXT UNIQUE NOT NULL,
                                     "fullname" CITEXT NOT NULL,
                                     "about"    TEXT
);

CREATE TABLE IF NOT EXISTS forums (
                                      "posts"   BIGINT  DEFAULT 0,
                                      "slug"    CITEXT  UNIQUE NOT NULL,
                                      "threads" INTEGER DEFAULT 0,
                                      "title"   TEXT    NOT NULL,
                                      "user"    CITEXT  NOT NULL REFERENCES users ("nickname")
);

CREATE TABLE IF NOT EXISTS threads (
                                       "id"      SERIAL         UNIQUE PRIMARY KEY,
                                       "author"  CITEXT         NOT NULL REFERENCES users ("nickname"),
                                       "created" TIMESTAMPTZ(3) DEFAULT now(),
                                       "forum"   CITEXT         NOT NULL REFERENCES forums ("slug"),
                                       "message" TEXT           NOT NULL,
                                       "slug"    CITEXT,
                                       "title"   TEXT           NOT NULL,
                                       "votes"   INTEGER        DEFAULT 0
);

CREATE TABLE IF NOT EXISTS posts (
                                     "id"       BIGSERIAL         UNIQUE PRIMARY KEY,
                                     "author"   CITEXT         NOT NULL REFERENCES users ("nickname"),
                                     "created"  TIMESTAMPTZ(3) DEFAULT now(),
                                     "forum"    CITEXT         NOT NULL REFERENCES forums ("slug"),
                                     "isEdited" BOOLEAN        DEFAULT FALSE,
                                     "message"  TEXT           NOT NULL,
                                     "parent"   INTEGER        DEFAULT 0,
                                     "thread"   INTEGER        NOT NULL REFERENCES threads ("id"),
                                     "path"     BIGINT []
);

CREATE TABLE IF NOT EXISTS votes (
                                     "thread"   INT NOT NULL REFERENCES threads("id"),
                                     "voice"    INTEGER NOT NULL,
                                     "nickname" CITEXT   NOT NULL
);


CREATE TABLE forum_users
(
    "forum_user"  CITEXT COLLATE ucs_basic NOT NULL,
    "forum"       CITEXT NOT NULL
);
`

type DbManager interface {
	DbConnect() error
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

const psqlURI = "postgresql://forum:forum@localhost:5432/mydb"

func (db *InitDB) DbConnect() (err error) {
	l := logrus.New()
	l.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	//logger := logrusadapter.NewLogger(l)
	runtimeParams := make(map[string]string)
	runtimeParams["application_name"] = "Web application"
	conConfig, _ := pgx.ParseURI(psqlURI)
	/*	pgx.ConnConfig{
		Host:           "127.0.0.1",
		Port:           5432,
		Database:       "forum",
		User:           "forum",
		Password:       "forum",
		TLSConfig:      nil,
		UseFallbackTLS: false,
		RuntimeParams:  runtimeParams,
		LogLevel:       5,
		Logger:         logger,
	}*/

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
	_, err = db.pool.Exec(sqlInit)
	if err != nil {
		return err
	}
	return nil
}
