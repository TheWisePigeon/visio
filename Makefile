.PHONY: setup-db stop-db delete-db reboot-db

build:
	@go build -o server ./main.go

run-dev:
	@go run main.go

install-deps:
	@go mod download

stop-db:
	@docker stop postgres

delete-db:
	@docker rm postgres

reboot-db: stop-db delete-db setup-db

migrate:
	@docker cp ./schema.sql postgres:/tmp/schema.sql
	@docker exec -it postgres psql -U postgres -d visio -f /tmp/schema.sql
	@echo "Migration successful"

build-image:
	@docker build -t visio:latest .

push-image:
	@docker tag visio:latest thewisepigeon/visio:latest
	@docker push thewisepigeon/visio:latest

start-tailwind-compilation:
	@npx tailwindcss -i ./assets/app.css -o ./public/output.css --minify --watch

build-css:
	@npx tailwindcss -i ./assets/app.css -o ./public/output.css --minify

test:
	@go test -p 1 -v ./...
