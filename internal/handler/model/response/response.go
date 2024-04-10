package response

import (
	"encoding/json"
	"time"
)

type Banner struct {
	Id        uint64
	TagsId    []int64
	FeatureId uint64
	Content   json.RawMessage
	IsActive  bool
	Created   time.Time
	Updated   time.Time
}
