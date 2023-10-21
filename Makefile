createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todolistwebapi
	
dropdb:
	docker exec -it postgres12 dropdb todolistwebapi


.PHONY: createdb dropdb