package crawlerraw

import "github.com/kevin-zx/wdtool/crawler-raw/dbentities"

type Repository interface {
	SaveRaw(tableName string, datas []*dbentities.CrawlerRawData) error
	Close() error
}
