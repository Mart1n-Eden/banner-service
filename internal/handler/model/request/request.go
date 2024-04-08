package request

import "encoding/json"

type PostBanner struct {
	TagIds    []uint64        `json:"tag_ids"`
	FeatureId uint64          `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}
