package dbentities

import "time"

type CrawlerRawData struct {
	Seed      string    `json:"seed"       db:"seed"`
	Data      string    `json:"data"       db:"data"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
