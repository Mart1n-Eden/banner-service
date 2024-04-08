package request

type PostBanner struct {
	TagIds    []uint64 `json:"tag_ids"`
	FeatureId uint64   `json:"feature_id"`
	Content   string   `json:"content"`
	IsActive  bool     `json:"is_active"`
}
