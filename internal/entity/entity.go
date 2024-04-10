package entity

type Banner struct {
	TagId     uint64 `json:"tag_id"`
	FeatureId uint64 `json:"feature_id"`
	Content   string `json:"content"`
	//IsActive bool
}
