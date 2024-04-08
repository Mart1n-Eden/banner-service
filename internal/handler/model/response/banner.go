package model

import "time"

type Banner struct {
	Id uint64
	//Feature  uint64
	//Tags     uint64
	Content  string
	IsActive bool
	Created  time.Time
	Updated  time.Time
}
