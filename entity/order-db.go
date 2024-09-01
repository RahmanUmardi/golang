package entity

import (
	"database/sql"
	"time"
)

type Order struct {
	Order_id        int
	Customer_id     int
	Order_date      time.Time
	Completion_date sql.NullTime
	Received_by     string
	Created_at      time.Time
	Updated_at      time.Time
}
