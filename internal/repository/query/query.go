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
			FROM unnest($3::bigint[]) as tag`

	DeleteBanner = `DELETE FROM banners
			WHERE id = $1`

	PatchContent = `UPDATE banners
			SET content = $1, updated_at = NOW()
			WHERE id = $2`

	PatchIsActive = `UPDATE banners
			SET is_active = $1, updated_at = NOW()
			WHERE id = $2`

	PatchFeature = `UPDATE banner_identifier
			SET feature_id = $1
			WHERE banner_id = $2`

	SelectFeature = `SELECT feature_id
			FROM banner_identifier
			WHERE banner_id = $1`

	DeleteIdentifier = `DELETE FROM banner_identifier
			WHERE banner_id = $1`
)
