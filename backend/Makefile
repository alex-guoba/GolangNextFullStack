DBUSER=root
DBSECRET=helloworld
DBNAME=blog_service
DB_URL=mysql://$(DBUSER):$(DBSECRET)@localhost:3306/$(DBNAME)?sslmode=disable

VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

.DEFAULT_GOAL := lint

# build
build:
	go build -o ./bin/gin-clean-template ./cmd/main.go

clean:
	rm ./bin/gin-clean-template

# Start mysql container
mysql_install:
	docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=$(DBSECRET) mysql

# Create database
db_create:
	docker exec -it mysql-test mysql -p$(DBSECRET) -e "create database $(DBNAME)"

# Generate new migration
migration_new:
	migrate create -ext sql -dir db/migration -seq $(name)

# lint source file
lint:
	golangci-lint run

# unit test coverage
unittest:
	go test ./... --covermode=count '-gcflags=all=-N -l' -v

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag/gen@latest
	go get github.com/swaggo/swag/cmd/swag@latest
	cd ./server && $$(go env GOPATH)/bin/swag init --parseDependency -g server.go

.PHONY: mysql_install db_create lint
