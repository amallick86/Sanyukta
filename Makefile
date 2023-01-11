DB_URL = postgres://root:root123@localhost:5432/sanyukta?sslmode=disable
dcup:
		sudo docker-compose up
dcdb:
		sudo docker-compose up postgres -d
dcdown:
		sudo docker-compose down
drmi:
		sudo docker rmi sanyukta-1
swagger:
		swag init
migcrt:
		goose -dir ./db/migration/ create table_name sql
gooseup:
		goose -dir ./db/migration/ -v postgres "$(DB_URL)" up
goosedown:
		goose -dir ./db/migration/ -v postgres "$(DB_URL)" down
dbdocs:
		dbdocs build docs/db.dbml
dbml2sql:
		dbml2sql --postgres -o docs/schema.sql docs/db.dbml

.PHONY: dcup dcdb dcdown drmi swagger migcrt gooseup goosedown dbdocs dbml2sql