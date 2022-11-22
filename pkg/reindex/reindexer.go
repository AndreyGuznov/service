package reindex

import (
	"fmt"
	"serv/pkg/logger"
	"sync"

	"github.com/restream/reindexer"

	_ "github.com/restream/reindexer/bindings/cproto"
)

var (
	conn   *RDB
	connMx = &sync.Mutex{}
)

type RDB struct {
	Instance *reindexer.Reindexer
}

func Conn() *RDB {
	connMx.Lock()
	defer connMx.Unlock()

	if conn == nil {
		var err error

		conn, err = Init()
		if err != nil {
			logger.Err("Error of initialization Reindexer", err)
			return nil
		}
	}

	return conn
}

func doInit() (*RDB, error) {
	addr := fmt.Sprintf("cproto://%s/%s", "127.0.0.1:6534", "DB")

	instance := reindexer.NewReindex(addr, reindexer.WithCreateDBIfMissing())

	db := RDB{
		Instance: instance,
	}

	return &db, nil
}

func Init() (*RDB, error) {
	var dbsess *RDB
	var err error

	dbsess, err = doInit()
	if err != nil {
		logger.Err("Failed to connect to Reindexer", err)
		return nil, err
	}

	logger.Info("Connected to Reindexer database")
	return dbsess, nil
}

func (db *RDB) Close() error {
	return db.Close()
}
