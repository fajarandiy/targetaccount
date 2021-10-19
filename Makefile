mysql:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_USER=adminDB -e MYSQL_PASSWORD=simaS123 -e MYSQL_ROOT_PASSWORD=simaS123 -e MYSQL_DATABASE=targetaccountdb -d mysql:8.0.26

mysqlWithCopyData:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_USER=adminDB -e MYSQL_PASSWORD=simaS123 -e MYSQL_ROOT_PASSWORD=simaS123 -e MYSQL_DATABASE=targetaccountdb -v /Users/fajarandiyulianto/Database/MySQL:/var/lib/mysql -d mysql:8.0.26

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "mysql://adminDB:simaS123@tcp(localhost:3306)/targetaccountdb" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://adminDB:simaS123@tcp(localhost:3306)/targetaccountdb" -verbose down

test:
	go test -v -cover  ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/Store.go com.example/targetaccount/db/sqlc Store

.PHONY: mysql mysqlWithCopyData createdb dropdb sqlc migrateup migratedown test server mock