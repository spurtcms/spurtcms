run:
	go run main.go

buildwithview:
	mkdir build
	cp -r lang build
	cp -r locales build
	cp -r public build
	cp -r view build
	cp -r cms.sql build
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o build/spurtcms-admin

build:
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o build/spurtcms-admin

permission:
	chmod -R 777 spurtcms-admin

start:
	sudo systemctl start spurtcms-admin