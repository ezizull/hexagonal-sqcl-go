CREATE TABLE IF NOT EXISTS todos (
	id BIGSERIAL PRIMARY KEY,
	activity_group_id int REFERENCES activities(id),
	title VARCHAR,
	is_active BOOLEAN,
	"priority" VARCHAR(120),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);
