header := Authorization: Bearer $(ACCESS_TOKEN)
host := http://localhost:12345


post_login:
	@curl -X POST -d '{"name": "john"}' $(host)/login | jq


get_subscriptions:
	@curl -H '$(header)' $(host)/subscriptions | jq


patch_subscription:
	@curl -X PATCH -H '$(header)' $(host)/subscriptions/$(SUBSCRIPTION_ID) -d '{"quantity": 14}' | jq


delete_subscription:
	@curl -X DELETE -H '$(header)' $(host)/subscriptions/$(SUBSCRIPTION_ID) | jq


create_subscription:
	@curl -X POST -H '$(header)' $(host)/subscriptions -d '{"productId": 1, "quantity": 14}' | jq
