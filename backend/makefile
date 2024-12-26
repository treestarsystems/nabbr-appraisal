BINARY_NAME=nabbr-appraisal-tool-backend-api.service

build:
	mkdir -p ../dist/backend
	go build -o ../dist/backend/${BINARY_NAME} main.go
	cp .env ../dist/backend/.env

run:
	mkdir -p ../dist/backend
	go build -o ../dist/backend/${BINARY_NAME} main.go
	cp .env ../dist/backend/.env
	cd ../dist/backend && nohup ./${BINARY_NAME} &

clean:
	go clean -cache
	rm ../dist/backend/${BINARY_NAME}