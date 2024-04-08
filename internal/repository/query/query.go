package query

const (
	GetBanner = `SELECT content 
			FROM banners 
			WHERE id in (
			SELECT banner_id 
			FROM banner_identifier 
			WHERE tag_id = $1 AND feature_id = $2) AND is_active 
			LIMIT 1`

	PostBanner = `INSERT INTO banners (content, is_active)
			VALUES ($1, $2)
			RETURNING id`

	PostIdentifiers = `INSERT INTO banner_identifier (banner_id, feature_id, tag_id)
			SELECT $1, $2, tag
			FROM unnest($3::bigint[]) as tag
`
)
