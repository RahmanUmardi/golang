package entity

import "time"

type Service struct {
	Service_id   int
	Service_name string
	Unit         string
	Price        int
	Created_at   time.Time
	Updated_at   time.Time
}
