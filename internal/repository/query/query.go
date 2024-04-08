package query

const (
	GetUser = `SELECT content 
			FROM banners 
			WHERE id in (
			SELECT banner_id 
			FROM banner_identifier 
			WHERE tag_id = $1 AND feature_id = $2) AND is_active 
			LIMIT 1`
)
