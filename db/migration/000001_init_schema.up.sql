CREATE TABLE tasks (
	id serial PRIMARY KEY NOT NULL,
	title varchar NOT NULL,
	"description" varchar NOT NULL,
	complete varchar NOT NULL
	-- created_at timestamp NOT NULL,
	-- updated_at timestamp NOT NULL,
	-- deleted_at timestamp NOT NULL
)