CREATE TABLE restock_notification_subscriptions (
	id uuid DEFAULT gen_random_uuid(),

	-- Foreign keys.
	product_id bigint NOT NULL,
	user_id uuid NOT NULL,
	quantity int NOT NULL CHECK (quantity > 0),

	-- Timestamps.
	created_at timestamptz NOT NULL DEFAULT current_timestamp,
	updated_at timestamptz NOT NULL DEFAULT current_timestamp,

	-- Constraints.
	PRIMARY KEY (id),
	FOREIGN KEY (product_id) REFERENCES products(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TRIGGER restock_notification_subscription_moddatetime
	BEFORE UPDATE ON restock_notification_subscriptions
	FOR EACH ROW
	EXECUTE PROCEDURE moddatetime(updated_at);
