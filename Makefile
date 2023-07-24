run-srv:
	go run ./cmd/app/main.go

run-db:
	docker run -d -p 27017:27017 --name mongo_db mongo

into-mongo:
	docker exec -it mongo_db mongosh
	# "show dbs" - display all databases
	# "use url-db" - switch to database
	# "show collections" - display all collections (tables)