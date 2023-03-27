CREATE TABLE IF NOT EXISTS activities (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR,
	email VARCHAR(320),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);
