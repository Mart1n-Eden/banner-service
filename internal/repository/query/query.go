package query

const (
	GetBanner = `SELECT content, is_active 
			FROM banners 
			WHERE id in (
			SELECT banner_id 
			FROM banner_identifier 
			WHERE tag_id = $1 AND feature_id = $2) AND is_active 
			LIMIT 1`

	PostBanner = `INSERT INTO banners (content, is_active)
			VALUES ($1, $2)
			RETURNING id`

	GetAdminBanner = `WITH tmp as (
			SELECT banner_id, feature_id, array_agg(tag_id)::bigint[] AS tag_ids 
			FROM banner_identifier
			WHERE banner_id in (
							SELECT DISTINCT banner_id 
							FROM banner_identifier 
							WHERE ($1::bigint IS NULL OR tag_id = $1) AND ($2::bigint IS NULL OR feature_id = $2))
			GROUP BY banner_id, feature_id)
			SELECT id, tag_ids, feature_id, content, is_active, created_at, updated_at 
			FROM banners 
			INNER JOIN tmp ON (tmp.banner_id = banners.id)
			LIMIT $3 OFFSET $4`

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
