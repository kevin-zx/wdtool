package crawlerraw

import "github.com/kevin-zx/wdtool/crawlerraw/dbentities"

type Repository interface {
	SaveRaw(tableName string, datas []*dbentities.CrawlerRawData) error
	Close() error
}
