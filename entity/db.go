package entity

import "time"

type DbEntity struct {
	ID           int       `db:"id"`
	CreatedBy    string    `db:"create_by"`
	CreatedTime  time.Time `db:"create_time"`
	ModifiedBy   string    `db:"create_by"`
	ModifiedTime time.Time `db:"modify_time"`
}

func (e *DbEntity) Insert(user string) {
	e.CreatedTime = time.Now()
	e.CreatedBy = user
}

func (e *DbEntity) Update(user string) {
	e.ModifiedBy = user
	e.ModifiedTime = time.Now()
}
