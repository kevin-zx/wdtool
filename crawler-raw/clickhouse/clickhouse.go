package clickhouse

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kevin-zx/kbase/kclickhouse"
	crawlerraw "github.com/kevin-zx/wdtool/crawler-raw"
	"github.com/kevin-zx/wdtool/crawler-raw/dbentities"
)

type clickhouseRepository struct {
	db *sqlx.DB
}

func NewClickhouseRepository(
	host string,
	port int,
	username string,
	password string,
	databese string,
) (crawlerraw.Repository, error) {
	db, err := kclickhouse.CreateClickhouseDB(host, port, username, password, databese)
	if err != nil {
		return nil, err
	}

	return &clickhouseRepository{
		db: db,
	}, nil
}

func (c *clickhouseRepository) Close() error {
	return c.db.Close()
}

func (c *clickhouseRepository) SaveRaw(tableName string, datas []*dbentities.CrawlerRawData) error {
	preparedSql := fmt.Sprintf(`
	INSERT INTO %s (seed, data, create_time)
	VALUES
	(:seed, :data, :create_time)
	`, tableName)
	tx := c.db.MustBegin()
	stmt, err := tx.PrepareNamed(preparedSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, data := range datas {
		if data == nil {
			continue
		}
		_, err := stmt.Exec(data)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
