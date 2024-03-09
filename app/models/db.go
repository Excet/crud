package models

type Graphs struct {
	ID     int64  `sql:"not null;unique"`
	Name   string `sql:"type:varchar(60)"`
	Parent int64  `sql:"not null"`
}
