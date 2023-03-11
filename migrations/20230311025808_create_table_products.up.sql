CREATE TABLE products (
	id bigint GENERATED ALWAYS AS IDENTITY,

	-- Attributes.
	name text NOT NULL,
	quantity int NOT NULL CHECK (quantity >= 0),

	-- Timestamps.
	created_at timestamptz NOT NULL DEFAULT current_timestamp,
	updated_at timestamptz NOT NULL DEFAULT current_timestamp,

	-- Constraints.
	PRIMARY KEY (id)
);

CREATE TRIGGER products_moddatetime
	BEFORE UPDATE ON products
	FOR EACH ROW
	EXECUTE PROCEDURE moddatetime(updated_at);
