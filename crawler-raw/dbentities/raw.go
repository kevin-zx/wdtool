package dbentities

import "time"

type CrawlerRawData struct {
	Seed       string    `json:"seed"        db:"seed"`
	Data       string    `json:"data"        db:"data"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
