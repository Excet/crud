package models

type GraphNode struct {
	ID     int32  `db:"id"`
	Type   string `db:"type"`
	Name   string `db:"name"`
	Parent int32  `db:"parent"`
	Price  int32  `db:"price"`
}
