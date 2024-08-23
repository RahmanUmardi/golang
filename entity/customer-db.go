package entity

import "time"

type Customer struct {
	Customer_id int
	Name        string
	Phone       int
	Address     string
	Updated_at  time.Time
	Created_at  time.Time
}
