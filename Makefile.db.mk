pg_conn_str := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate := docker run \
					 --rm \
					 --network=host  \
					 -v $(PWD)/migrations:/migrations \
					 migrate/migrate \
					 -path=migrations \
					 -database $(pg_conn_str)


# Display migrate commands.
help:
	@$(migrate) -help


# Create new migration file. e.g. make create name=create_table_users
create:
	@$(migrate) create -ext=sql -dir=migrations $(name)


# Run all migrations.
migrate:
	@$(migrate) up


# Rollback a migration. Use n=-all for full rollback.
# n defaults to 1.
rollback:
	@$(migrate) down $${n:=1}


# Drop everything in the database.
drop:
	@$(migrate) drop -f
