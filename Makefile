createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todolistwebapi

dropdb:
	docker exec -it postgres12 dropdb todolistwebapi

create-migration:
	migrate create -ext sql -dir database/migrations -seq todolist_schema

migrateup:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" -verbose down

atlas-migration:
	atlas migrate diff --env gorm 

atlas-push-migration:
	atlas migrate push app --dev-url "docker://postgres/15/dev?search_path=public"

atlas-apply-migration:
	atlas migrate apply --dir "atlas://app"  --url "postgres://root:secret@:5432/todolistwebapi?search_path=public&sslmode=disable"

.PHONY: createdb dropdb atlas-migration atlas-apply-migration


