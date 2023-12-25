ifneq (,$(wildcard ./.env))
  include .env
  export
endif

migrate-up:
	@migrate -path migrations -database "${PG_URL}" up

migrate-down:
	@migrate -path migrations -database "${PG_URL}" down

migrate-reset: migrate-down migrate-up

run:
	GIN_MODE=debug go run main.go