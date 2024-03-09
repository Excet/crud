package models

type Graphs struct {
	ID       int64   `db:"id"`
	Name     string  `db:"name"`
	Children []int64 `db:"children"`
	Parent   int64   `db:"parent"`
}
