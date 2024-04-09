package request

import "encoding/json"

type Banner struct {
	TagIds    []uint64         `json:"tag_ids,omitempty"`
	FeatureId *uint64          `json:"feature_id,omitempty"`
	Content   *json.RawMessage `json:"content,omitempty"`
	IsActive  *bool            `json:"is_active,omitempty"`
}
