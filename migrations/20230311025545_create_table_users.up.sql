CREATE EXTENSION moddatetime;


CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid(),

	-- Attributes.
	name text NOT NULL,

	-- Timestamps.
	created_at timestamptz NOT NULL DEFAULT current_timestamp,
	updated_at timestamptz NOT NULL DEFAULT current_timestamp,

	-- Constraints.
	PRIMARY KEY (id),
	UNIQUE (name)
);


CREATE TRIGGER user_moddatetime
	BEFORE UPDATE ON users
	FOR EACH ROW
	EXECUTE PROCEDURE moddatetime(updated_at);
