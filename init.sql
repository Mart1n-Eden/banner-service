CREATE TABLE IF NOT EXISTS banners
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    content JSONB NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS banner_identifier
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    banner_id BIGINT NOT NULL REFERENCES banners (id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL,
    feature_id BIGINT NOT NULL
);
-- WITH sb AS ( SELECT banner_id FROM banner_identifier WHERE tag_id = 1 and feature_id = 2 ) SELECT content FROM banners INNER JOIN sb on (sb.banner_id = banners.id) WHERE is_active LIMIT 1
-- SELECT content FROM banners WHERE id in (SELECT banner_id FROM banner_identifier WHERE tag_id = 1 and feature_id = 1) AND is_active LIMIT 1